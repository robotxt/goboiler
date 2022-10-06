package topic

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type TopicService interface {
}

type service struct {
	ctx context.Context
	log *logrus.Logger
	db  *gorm.DB
}

func NewTopicService(ctx context.Context, db *gorm.DB, logger *logrus.Logger) TopicService {
	return &service{
		ctx: ctx,
		log: logger,
		db:  db,
	}
}
