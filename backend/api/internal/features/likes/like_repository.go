package likes

import (
	"Server/internal/helpers"
	"Server/internal/store"
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type likeRepository struct {
	collection *mongo.Collection
}

func NewlikeRepository(db *store.Database) LikeRepository {
	col := db.Database.Collection("likes")

	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "postId", Value: 1}, {Key: "status", Value: 1}}},
		{Keys: bson.D{{Key: "userId", Value: 1}, {Key: "postId", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	col.Indexes().CreateMany(ctx, indexes)

	return &likeRepository{collection: col}
}

func (r *likeRepository) DeleteLike(like Like) error {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": like.ID}, bson.M{"$set": bson.M{
		"status": 0,
	}})
	if err != nil {
		return errors.New("cannot delete")
	}
	if result.ModifiedCount == 0 {
		return errors.New("no document deleted")
	}
	return nil
}

func (r *likeRepository) AddLike(like Like) error {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	if hasLike, err := r.hasLikeUnscoped(like.PostID, like.UserID); err != nil {
		return err
	} else if hasLike {
		return r.LikePost(like)
	}
	_, err := r.collection.InsertOne(ctx, like)
	if err != nil {
		return errors.New("cannot add like")
	}
	return nil
}

func (r *likeRepository) LikePost(like Like) error {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	result, err := r.collection.UpdateOne(ctx, bson.M{"userId": like.UserID, "postId": like.PostID}, bson.M{"$set": bson.M{
		"status": 1,
	}})
	if err != nil {
		return errors.New("cannot like post")
	}
	if result.ModifiedCount == 0 {
		return errors.New("no document updated")
	}

	return nil
}

func (r *likeRepository) DeleteLikesFromPost(postId bson.ObjectID) error {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	_, err := r.collection.UpdateMany(ctx, bson.M{"postId": postId}, bson.M{"$set": bson.M{
		"status": 0,
	}})
	if err != nil {
		return errors.New("cannot delete likes from post")
	}
	return nil
}

func (r *likeRepository) HasLike(postId, userId bson.ObjectID) (bool, error) {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"postId": postId, "userId": userId, "status": 1})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *likeRepository) hasLikeUnscoped(postId, userId bson.ObjectID) (bool, error) {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"postId": postId, "userId": userId})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *likeRepository) GetLikesCountByPostId(postId bson.ObjectID) (int64, error) {
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"postId": postId, "status": 1})
	if err != nil {
		return 0, errors.New("cannot get likes by post id")
	}

	return count, nil
}
