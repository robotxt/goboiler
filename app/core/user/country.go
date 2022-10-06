package user

import (
	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	UID       string `gorm:"unique" json:"uid"`
	Name      string `gorm:"type:varchar(50)" json:"name"`
	Currency  string `gorm:"type:varchar(20);default:'php'" json:"currency"`
	Tag       string `gorm:"type:varchar(20)" json:"tag"`
	ShortName string `gorm:"type:varchar(20)" json:"short_name"`
	Timezone  string `gorm:"type:varchar(20)" json:"timezone"`
}

func (srv service) DefaultCountry() (*Country, error) {
	default_country := "philippines"

	country := Country{}
	srv.db.Where(Country{Tag: default_country}).Find(&country)

	return &country, nil
}
