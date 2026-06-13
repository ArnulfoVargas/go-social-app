package posts

import (
	"Server/internal/features/media"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Post struct {
	ID            bson.ObjectID `bson:"_id" json:"id"`
	UserID        bson.ObjectID `bson:"userId" json:"userId"`
	Content       string        `bson:"content,omitempty" json:"content,omitempty"`
	Status        uint8         `bson:"status,default=1" json:"status,omitempty"`
	Media         []media.Media `bson:"media,omitempty" json:"media,omitempty"`
	LikesCount    int           `bson:"likesCount" json:"likesCount"`
	CommentsCount int           `bson:"commentsCount" json:"commentsCount"`
	CreatedAt     bson.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt     bson.DateTime `bson:"updatedAt" json:"updatedAt"`
}
