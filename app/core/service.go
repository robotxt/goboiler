package core

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/robotxt/goboiler/app/core/topic"
	"github.com/robotxt/goboiler/app/core/user"
	"github.com/sirupsen/logrus"
)

type Service interface {
	BasicUserRegistration(email string, pwd string) (*user.User, error)
	BasicLoginAuth(email string, pwd string) (*user.Token, error)
	ValidateUserToken(token_str string) (*user.User, error)
}

type service struct {
	ctx      context.Context
	log      *logrus.Logger
	userSrv  user.UserService
	topicSrv topic.TopicService
}

func NewCoreService(ctx context.Context, db *gorm.DB, logger *logrus.Logger) Service {
	user_service := user.NewUserService(ctx, db, logger)
	topic_service := topic.NewTopicService(ctx, db, logger)

	// Migrations
	user_service.UserMigration()

	return &service{
		ctx:      ctx,
		log:      logger,
		userSrv:  user_service,
		topicSrv: topic_service,
	}
}
