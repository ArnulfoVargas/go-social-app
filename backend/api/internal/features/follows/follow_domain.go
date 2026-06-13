package follows

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type FollowService interface {
	ToggleFollowUser(userID, targetUserID string) (bool, error)
	GetFollowingCount(userID string) (int64, error)
	GetFollowersCount(userID string) (int64, error)
}

type FollowRepository interface {
	FollowUser(follow Follow) error
	UnfollowUser(userID, targetUserID bson.ObjectID) error
	UserIsFollowing(userID, targetUserID bson.ObjectID) (bool, error)
	GetFollowingCount(userID bson.ObjectID) (int64, error)
	GetFollowersCount(userID bson.ObjectID) (int64, error)
	GetFollowingIds(userID bson.ObjectID) ([]bson.ObjectID, error)
	GetRelatedFollowSuggestions(userId bson.ObjectID, followingIds []bson.ObjectID, limit int) ([]bson.ObjectID, error)
}
