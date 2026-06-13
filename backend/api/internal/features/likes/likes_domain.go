package likes

import "go.mongodb.org/mongo-driver/v2/bson"

type LikeRepository interface {
	DeleteLike(like Like) error
	AddLike(like Like) error
	HasLike(postId, userId bson.ObjectID) (bool, error)
	DeleteLikesFromPost(postId bson.ObjectID) error
	GetLikesCountByPostId(postId bson.ObjectID) (int64, error)
}
