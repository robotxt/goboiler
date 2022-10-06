package user

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func CreateDefaultUserGroup(srv service) *gorm.DB {
	group := Group{}
	srv.db.Where(Group{Name: "admin"}).Find(&group)

	if group.UID == "" {
		// add default admin group
		admin_uid := uuid.New().String()

		admin_group := Group{}
		admin_group.UID = admin_uid
		admin_group.Name = "admin"
		srv.db.Create(&admin_group)
	}

	user_group := Group{}
	srv.db.Where(Group{Name: "user"}).Find(&user_group)
	if user_group.UID == "" {
		// add default user group
		user_uid := uuid.New().String()

		new_user_group := Group{}
		new_user_group.UID = user_uid
		new_user_group.Name = "user"
		srv.db.Create(&new_user_group)
	}

	return srv.db
}

func (srv service) UserMigration() *gorm.DB {
	srv.log.Info("Applying Users Service DB migrations")

	srv.db.AutoMigrate(User{})
	srv.db.AutoMigrate(UserGroup{})
	srv.db.AutoMigrate(Country{})
	srv.db.AutoMigrate(Token{})
	srv.db.AutoMigrate(Group{})

	def_country := "philippines"

	country := Country{}
	srv.db.Where(Country{Tag: def_country}).Find(&country)
	if country.UID == "" {
		// add default country if not exist
		uid := uuid.New().String()

		new_country := Country{}
		new_country.UID = uid
		new_country.Name = def_country
		new_country.Currency = "php"
		new_country.Tag = def_country
		new_country.ShortName = "ph"
		new_country.Timezone = "Asia/Manila"
		srv.db.Create(&new_country)

	}

	// create default user groups
	CreateDefaultUserGroup(srv)
	return srv.db
}
