package image

import "github.com/jinzhu/gorm"

type Image struct {
	gorm.Model

	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Fullsize  string `json:"fullsize"`
	UID       string `gorm:"unique" json:"uid"`
}
