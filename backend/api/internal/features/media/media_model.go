package media

import "go.mongodb.org/mongo-driver/v2/bson"

type Media struct {
	ID       bson.ObjectID `bson:"_id" json:"id"`
	URL      string        `bson:"url" json:"url"`
	PublicID string        `bson:"publicId" json:"publicId"`
}
