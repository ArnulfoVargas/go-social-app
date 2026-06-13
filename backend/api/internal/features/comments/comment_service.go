package comments

type commentsService struct {
	commentsRepo CommentRepository
}

func NewCommensService(cr CommentRepository) CommentService {
	return &commentsService{
		commentsRepo: cr,
	}
}

func (cs commentsService) AddComment(postId, userId, content string) error {
	return nil
}

func (cs commentsService) GetComments(postId string, limit int) ([]Comment, error) {
	return nil, nil
}

func (cs commentsService) GetCommentsCountById(postId string) (int64, error) {
	return 0, nil
}

func (cs commentsService) DeleteComment(commentId, userId string) error {
	return nil
}

func (cs commentsService) UpdateComment(commentId, userId, content string) error {
	return nil
}
