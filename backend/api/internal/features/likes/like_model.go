package likes

import "go.mongodb.org/mongo-driver/v2/bson"

type Like struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	PostID    bson.ObjectID `bson:"postId" json:"postId" validate:"required"`
	UserID    bson.ObjectID `bson:"userId" json:"userId" validate:"required"`
	Status    uint8         `bson:"status" json:"status" validate:"required"`
	CreatedAt bson.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt bson.DateTime `bson:"updatedAt" json:"updatedAt"`
}
