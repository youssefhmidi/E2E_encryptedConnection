package models

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	PublicKey      string `gorm:"unique"`
	Name           string
	Email          string `gorm:"unique"`
	Pasword        string
	NickName       string
	Messages       []Message  `gorm:"foreignKey:UserID"`
	ChatRoomsOwned []ChatRoom `gorm:"foreignKey:OwnerID"`
	ChatRooms      []ChatRoom `gorm:"many2many:user_chatroom"`
}

type UserRepository interface {
	//	Create a user
	CreatUser(ctx context.Context, user User) error

	//	get a user

	GetUserById(ctx context.Context, ID uint) (User, error)
	GetUserByPublicKey(ctx context.Context, publicKey string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	//	fetching all users
	FetchAllUsers(ctx context.Context, limit int) ([]User, error)

	//	Update Information about a User

	//	     // target : a table column like "email, password, ...etc", E.g:
	//	  	UpdateUser(ctx, user, "name", "NewName")
	UpdateUser(ctx context.Context, user User, target string, value interface{}) error

	// delete User

	DeleteUser(ctx context.Context, user User) error
}
