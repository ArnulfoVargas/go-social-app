package feed

import (
	"Server/internal/constants"
	"Server/internal/features/comments"
	"Server/internal/features/media"
	"Server/internal/features/posts"
	"Server/internal/features/users"
	"Server/internal/helpers"
	"Server/internal/shared"
	"Server/internal/validator"
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/gofiber/fiber/v3"
)

type FeedHanler struct {
	postService     posts.PostService
	commentsService comments.CommentService
	userService     users.UserService
}

func NewFeedHandler(validator *validator.Validator, postService posts.PostService, commentsService comments.CommentService, userService users.UserService) *FeedHanler {
	return &FeedHanler{
		postService:     postService,
		commentsService: commentsService,
		userService:     userService,
	}
}

func SetupPostRoutes(s fiber.Router, postHandler *FeedHanler) {
	g := s.Group("/posts", shared.Protected(shared.ParseJWT))

	g.Get("/suggested", postHandler.getSuggestedPosts)
}

func (f *FeedHanler) getSuggestedPosts(c fiber.Ctx) error {
	id, ok := helpers.GetUserIdFromLocals(c)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(shared.ErrorResponse{
			Error:  "unauthorized",
			Status: fiber.StatusUnauthorized,
		})
	}

	postsSug, err := f.postService.GetSuggestedPosts(id, constants.DEFAULT_PAGE_SIZE)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(shared.ErrorResponse{
			Error:  "failed to get suggested posts",
			Status: fiber.StatusInternalServerError,
		})
	}

	suggestions := make([]PostSuggestion, len(postsSug))
	mu := sync.Mutex{}
	g, ctx := errgroup.WithContext(c.Context())

	uc := &userCache{
		users: make(map[string]*users.User, 0),
	}

	for i, p := range postsSug {
		post := p
		index := i
		g.Go(func() error {
			sf := suggestionFill{
				array:     &suggestions,
				post:      &post,
				mu:        &mu,
				index:     index,
				userCache: uc,
			}

			return f.fillSuggestions(ctx, sf)
		})
	}

	if err := g.Wait(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(shared.ErrorResponse{
			Error:  "couldnt get gosts",
			Status: fiber.StatusInternalServerError,
		})
	}

	return c.JSON(shared.GenericResponse[[]PostSuggestion]{
		Data: suggestions,
	})
}

type suggestionFill struct {
	array     *[]PostSuggestion
	post      *posts.Post
	mu        *sync.Mutex
	index     int
	userCache *userCache
}

func (f *FeedHanler) fillSuggestions(ctx context.Context, sf suggestionFill) error {
	post := *sf.post
	userId := post.UserID.Hex()
	user, err := sf.userCache.Get(userId, f.userService.GetUser)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	if err != nil {
		return err
	}

	postId := post.ID.Hex()

	lCount, err := f.postService.GetLikesCountByPostId(postId)
	if err != nil {
		return err
	}

	cCount, err := f.commentsService.GetCommentsCountById(postId)
	if err != nil {
		return err
	}

	comments, err := f.commentsService.GetComments(postId, 2)
	if err != nil {
		return err
	}

	cr := make([]CommentResponse, len(comments))

	commentsMutex := sync.Mutex{}
	ng, ctx := errgroup.WithContext(ctx)

	for i, c := range comments {
		cm := c
		index := i

		ng.Go(func() error {
			fc := fillComments{
				mu:                &commentsMutex,
				commentsResponses: &cr,
				index:             index,
				comment:           cm,
				userCache:         sf.userCache,
			}

			return f.fillComments(ctx, fc)
		})
	}

	if err := ng.Wait(); err != nil {
		return err
	}

	var avatar *media.MediaResponse

	if len(user.Avatar.URL) > 0 {
		avatar = &media.MediaResponse{
			ID:  user.Avatar.ID.Hex(),
			URL: user.Avatar.URL,
		}
	}
	userMin := users.UserResponseMin{
		ID:     userId,
		Name:   user.Name,
		Avatar: avatar,
	}

	sug := PostSuggestion{
		PostId:       post.ID.Hex(),
		User:         userMin,
		LikeCount:    lCount,
		CommentCount: cCount,
		Comments:     cr,
	}

	sf.mu.Lock()
	(*sf.array)[sf.index] = sug
	sf.mu.Unlock()

	return nil
}

type fillComments struct {
	mu                *sync.Mutex
	commentsResponses *[]CommentResponse
	index             int
	comment           comments.Comment
	userCache         *userCache
}

func (f *FeedHanler) fillComments(ctx context.Context, fc fillComments) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	userId := fc.comment.UserID.Hex()
	user, err := fc.userCache.Get(userId, f.userService.GetUser)
	if err != nil {
		return err
	}

	var avatar *media.MediaResponse

	if len(user.Avatar.URL) > 0 {
		avatar = &media.MediaResponse{
			ID:  user.Avatar.ID.Hex(),
			URL: user.Avatar.URL,
		}
	}

	userMin := users.UserResponseMin{
		ID:     userId,
		Name:   user.Name,
		Avatar: avatar,
	}

	fc.mu.Lock()

	(*fc.commentsResponses)[fc.index] = CommentResponse{
		User: userMin,
		Text: fc.comment.Content,
	}

	fc.mu.Unlock()
	return nil
}
