package topic

import (
	"github.com/jinzhu/gorm"
	"github.com/robotxt/goboiler/app/core/user"
)

type Topic struct {
	gorm.Model
	UID     string `gorm:"unique" json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tag     string `json:"tag"` // string separated by comma and #
	Private bool   `gorm:"default:false" json:"private"`

	UserID uint
	User   user.User `gorm:"foreignKey:UserID"`
}
