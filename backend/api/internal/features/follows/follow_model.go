package follows

import "go.mongodb.org/mongo-driver/v2/bson"

type Follow struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FollowerID  bson.ObjectID `bson:"followerId" json:"followerId"`
	FollowingID bson.ObjectID `bson:"followingId" json:"followingId"`
	CreatedAt   bson.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt   bson.DateTime `json:"updatedAt" bson:"updatedAt"`
	Status      uint8         `json:"status" bson:"status"`
}
