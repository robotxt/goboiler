package user

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

var default_expiration time.Duration = 4

type Token struct {
	gorm.Model
	Token      string    `gorm:"unique" json:"token"`
	Expiration time.Time `json:"expiration"`

	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}

func (srv service) GenerateToken(user *User) *Token {
	token_str := uuid.New().String()
	expiryDate := time.Now().Local().Add(time.Hour * default_expiration)

	token := Token{}
	token.Token = token_str
	token.Expiration = expiryDate
	token.UserID = user.ID
	token.User = *user
	srv.db.Create(&token)

	return &token
}

func (srv service) GetToken(user *User) *Token {
	token := Token{}
	srv.db.Where(Token{UserID: user.ID}).Last(&token)

	return &token
}

func (srv service) GetUserToken(token_str string) (*Token, error) {
	token := Token{}
	result := srv.db.Where(Token{Token: token_str}).Last(&token)

	return &token, result.Error
}
