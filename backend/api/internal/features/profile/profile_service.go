package profile

import (
	"Server/internal/features/follows"
	"Server/internal/features/media"
	"Server/internal/features/posts"
	"Server/internal/features/users"
	"Server/internal/helpers"
	"errors"
	"sync"
)

type profileService struct {
	userService   users.UserRepository
	followService follows.FollowRepository
	postService   posts.PostRepository
}

func NewProfileService(userRepo users.UserRepository, followRepo follows.FollowRepository, postRepo posts.PostRepository) *profileService {
	return &profileService{
		userService:   userRepo,
		followService: followRepo,
		postService:   postRepo,
	}
}

func (s *profileService) GetProfile(userID string) (*Profile, error) {
	uid, err := helpers.ToObjectID(userID)
	if err != nil {
		return nil, err
	}

	var (
		wg                   sync.WaitGroup
		user                 *users.User
		followers, following int64 = 0, 0
		firstErr             error
		userPosts            = make([]posts.Post, 0)
		mu                   = &sync.Mutex{}
	)

	setErr := func(err error) {
		mu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		mu.Unlock()
	}

	wg.Go(func() {
		u, err := s.userService.GetUserById(uid)
		if err != nil {
			setErr(err)
			return
		}
		if u == nil {
			setErr(errors.New("user not found"))
		}

		mu.Lock()
		user = u
		mu.Unlock()
	})

	wg.Go(func() {
		uP, err := s.postService.GetPostsByUserId(uid)
		if err != nil {
			setErr(err)
			return
		}

		mu.Lock()
		userPosts = uP
		mu.Unlock()
	})

	wg.Go(func() {
		fCount, err := s.followService.GetFollowersCount(uid)
		if err != nil {
			setErr(err)
			return
		}
		mu.Lock()
		followers = fCount
		mu.Unlock()
	})

	wg.Go(func() {
		followingCount, err := s.followService.GetFollowingCount(uid)
		if err != nil {
			setErr(err)
			return
		}
		mu.Lock()
		following = followingCount
		mu.Unlock()
	})

	wg.Wait()

	if firstErr != nil {
		return nil, firstErr
	}

	profile := &Profile{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Posts:     posts.PostsFromModels(userPosts),
		Followers: followers,
		Following: following,
	}

	if user.Avatar.URL != "" {
		profile.Avatar = &media.MediaResponse{
			ID:  user.Avatar.ID.Hex(),
			URL: user.Avatar.URL,
		}
	}

	return profile, nil
}
