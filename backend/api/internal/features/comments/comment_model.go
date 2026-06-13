package comments

import "go.mongodb.org/mongo-driver/v2/bson"

type Comment struct {
	ID        bson.ObjectID `bson:"_id" json:"id"`
	PostID    bson.ObjectID `bson:"postId" json:"postId"`
	UserID    bson.ObjectID `bson:"userId" json:"userId"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt bson.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt bson.DateTime `bson:"updatedAt" json:"updatedAt"`
}
