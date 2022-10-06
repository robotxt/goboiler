package user

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	UserMigration() *gorm.DB
	CreateUser(email string, pwd string, country *Country) (*User, error)
	DefaultCountry() (*Country, error)
	GenerateToken(user *User) *Token
	GetToken(user *User) *Token
	GetUserByEmail(email string) (*User, error)
	GetUserByUID(uid string) (*User, error)
	GetUserByID(user_id uint) (*User, error)
	GetUserToken(token_str string) (*Token, error)
}

type service struct {
	ctx context.Context
	log *logrus.Logger
	db  *gorm.DB
}

func NewUserService(ctx context.Context, db *gorm.DB, logger *logrus.Logger) UserService {
	return &service{
		ctx: ctx,
		log: logger,
		db:  db,
	}
}
