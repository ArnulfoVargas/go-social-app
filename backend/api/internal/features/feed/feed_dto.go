package feed

import (
	"Server/internal/features/users"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostSuggestion struct {
	PostId       primitive.ObjectID    `json:"postId"`
	User         users.UserResponseMin `json:"user"`
	LikeCount    int                   `json:"likeCount"`
	CommentCount int                   `json:"commentCount"`
	Comments     []CommentResponse     `json:"comments"`
}

type CommentResponse struct {
	User users.UserResponseMin `json:"user"`
	Text string                `json:"text"`
}
