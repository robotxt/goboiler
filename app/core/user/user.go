package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/robotxt/goboiler/app/core/image"
)

type User struct {
	gorm.Model
	UID      string `gorm:"unique" json:"uid"`
	Username string `gorm:"unique;type:varchar(100)" json:"username"`
	Email    string `gorm:"unique;type:varchar(100)" json:"email"`
	Password string `json:"password"`

	CountryID uint
	Country   Country `gorm:"foreignKey:CountryID"`

	ImageID uint
	Image   image.Image `gorm:"foreignKey:ImageID"`
}

type UserGroup struct {
	gorm.Model
	UserID uint
	User   User `gorm:"foreignKey:UserID"`

	GroupID uint
	Group   Group `gorm:"foreignKey:GroupID"`
}

func (srv service) CreateUser(email string, pwd string, country *Country) (*User, error) {
	srv.log.Info("creating user: " + email)
	uid := uuid.New().String()

	new_user := User{}
	new_user.UID = uid
	new_user.Email = email
	new_user.Username = uid
	new_user.Password = pwd
	new_user.CountryID = country.ID
	new_user.Country = *country

	srv.db.Create(&new_user)
	return &new_user, nil
}

func (srv service) GetUserByEmail(email string) (*User, error) {
	user := User{}
	result := srv.db.Where(User{Email: email}).Find(&user)

	return &user, result.Error
}

func (srv service) GetUserByID(user_id uint) (*User, error) {
	user := User{}
	result := srv.db.Where(User{Model: gorm.Model{ID: user.ID}}).Last(&user)

	return &user, result.Error
}

func (srv service) GetUserByUID(uid string) (*User, error) {
	user := User{}
	result := srv.db.Where(User{UID: uid}).Last(&user)

	return &user, result.Error
}
