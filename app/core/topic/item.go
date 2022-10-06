package topic

import (
	"github.com/jinzhu/gorm"
	"github.com/robotxt/goboiler/app/core/image"
	"github.com/robotxt/goboiler/app/core/user"
)

type Item struct {
	gorm.Model
	UID         string `gorm:"unique" json:"uid"`
	Name        string `gorm:"type:varchar(300)" json:"name"`
	Description string `json:"description"`

	TopicID uint
	Topic   Topic `gorm:"foreignKey:TopicID"`
}

type ItemVote struct {
	gorm.Model
	UID       string `gorm:"unique" json:"uid"`
	Comment   string `json:"comment"`
	Anonymous bool   `gorm:"default:false" json:"anonymous"`

	TopicID uint
	Topic   Topic `gorm:"foreignKey:TopicID"`

	ItemID uint
	Item   Item `gorm:"foreignKey:ItemID"`

	UserID uint
	User   user.User `gorm:"foreignKey:UserID"`
}

type ItemImage struct {
	gorm.Model
	UID     string `gorm:"unique" json:"uid"`
	Deleted bool   `json:"deleted"`

	ImageID uint
	Image   image.Image `gorm:"foreignKey:ImageID"`
}
