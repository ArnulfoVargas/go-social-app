package comments

import (
	"Server/internal/helpers"
	"Server/internal/store"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type commentsRepository struct {
	collection *mongo.Collection
}

func NewCommentsRepository(db *store.Database) CommentRepository {
	col := db.Database.Collection("comments")

	indexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "postId", Value: 1}, {Key: "status", Value: 1}, {Key: "createdAt", Value: -1}}},
	}
	ctx, cancel := helpers.GenerateContext()
	defer cancel()

	col.Indexes().CreateMany(ctx, indexes)

	return &commentsRepository{
		collection: col,
	}
}

func (c commentsRepository) AddComment(comment Comment) error {
	return nil
}

func (c commentsRepository) GetComments(postId primitive.ObjectID, limit int) ([]Comment, error) {
	return nil, nil
}

func (c commentsRepository) GetCommentsCountById(postId primitive.ObjectID) (int64, error) {
	return 0, nil
}

func (c commentsRepository) DeleteComment(commentId primitive.ObjectID) error {
	return nil
}

func (c commentsRepository) UpdateComment(commentId primitive.ObjectID, object bson.M) error {
	return nil
}
