package posts

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PostService interface {
	CreatePost(userId string, post PostAdd) (Post, error)
	GetPost(postId string) (Post, error)
	DeletePost(postId string) error
	UpdatePost(postId string, content UpdatePostRequest) (Post, error)
	GetPostsByUserId(userId string) ([]Post, error)
	ToggleLike(postId string, userId string) error
	GetSuggestedPosts(userId string, limit int) ([]Post, error)
	GetLikesCountByPostId(postId string) (int64, error)
}

type PostRepository interface {
	CreatePost(post Post) error
	GetPost(postId bson.ObjectID) (Post, error)
	DeletePost(postId bson.ObjectID) error
	UpdatePost(postId bson.ObjectID, update bson.M) (Post, error)
	GetPostsByUserId(userId bson.ObjectID) ([]Post, error)
	ExistsById(postId bson.ObjectID) (bool, error)
	GetSuggestedPosts(userId bson.ObjectID, limit int) ([]Post, error)
}
