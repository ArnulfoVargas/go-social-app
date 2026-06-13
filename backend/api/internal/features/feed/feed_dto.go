package feed

import (
	"Server/internal/features/users"
)

type PostSuggestion struct {
	PostId       string                `json:"postId"`
	User         users.UserResponseMin `json:"user"`
	LikeCount    int64                 `json:"likeCount"`
	CommentCount int64                 `json:"commentCount"`
	Comments     []CommentResponse     `json:"comments"`
}

type CommentResponse struct {
	User users.UserResponseMin `json:"user"`
	Text string                `json:"text"`
}
