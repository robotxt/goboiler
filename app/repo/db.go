package repo

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

type Service interface {
	Connect() (*gorm.DB, error)
}

type service struct {
	log *logrus.Logger
}

func NewRepoService(logger *logrus.Logger) Service {
	return &service{
		log: logger,
	}
}

func (srv service) Connect() (*gorm.DB, error) {

	dbinfo := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	conn, err := gorm.Open("postgres", dbinfo)
	if err != nil {
		srv.log.Fatal(err)
		srv.log.Fatal("Could not connect database")
	}
	return conn, err
}
