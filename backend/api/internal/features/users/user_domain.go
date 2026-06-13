package users

import (
	"Server/internal/features/media"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserService interface {
	GetUser(id string) (*User, error)
	UpdateUser(id string, user *UpdateProfileRequest) error
	DeleteUser(id string) error
	AddProfilePicture(id string, media media.Media) error
	RemoveProfilePicture(id string) error
	ExistsUser(id string) (bool, error)
}

type UserRepository interface {
	GetUserById(id bson.ObjectID) (*User, error)
	UpdateUserById(id bson.ObjectID, data bson.M) error
	UserExistsById(id bson.ObjectID) (bool, error)
	GetUsersExcluding(excludeIDs []bson.ObjectID, limit int) ([]User, error)
	GetUsersByIds(ids []bson.ObjectID) ([]User, error)
	GetIdsExcluding(excludeIDs []bson.ObjectID, limit int) ([]bson.ObjectID, error)
	DeleteUserById(id bson.ObjectID) error
	SetProfilePicture(id bson.ObjectID, media media.Media) error
	RemoveProfilePicture(id bson.ObjectID) error
}
