package user

import "github.com/jinzhu/gorm"

type Group struct {
	gorm.Model
	UID  string `gorm:"unique" json:"uid"`
	Name string `gorm:"unique;varchar(30)" json:"name"`
}
