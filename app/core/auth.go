package core

import (
	"errors"
	"time"

	"github.com/robotxt/goboiler/app/core/user"
	"golang.org/x/crypto/bcrypt"
)

var invalid_cred = errors.New("Invalid Credentials")

func (srv service) BasicLoginAuth(email string, pwd string) (*user.Token, error) {
	user, err := srv.userSrv.GetUserByEmail(email)

	if err != nil {
		// email not exists
		return nil, invalid_cred
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password), []byte(pwd)); err != nil {
		// password not match
		return nil, invalid_cred
	}

	now_date := time.Now().Local()
	existing_token := srv.userSrv.GetToken(user)

	if existing_token.Expiration.Before(now_date) {
		// create new token if exisitng token expired
		token := srv.userSrv.GenerateToken(user)
		return token, nil
	}

	return existing_token, nil
}

func (srv service) ValidateUserToken(token_str string) (*user.User, error) {
	token, err := srv.userSrv.GetUserToken(token_str)
	if err != nil {
		return nil, invalid_cred
	}

	now_date := time.Now().Local()
	if token.Expiration.Before(now_date) {
		return nil, invalid_cred
	}

	user, _ := srv.userSrv.GetUserByID(token.UserID)

	return user, err
}
