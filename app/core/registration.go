package core

import (
	"errors"

	"github.com/robotxt/goboiler/app/core/user"
)

func (srv service) BasicUserRegistration(email string, pwd string) (*user.User, error) {
	hashed_pass, err := HashPassword([]byte(pwd))

	if err != nil {
		return nil, err
	}

	existing_user, _ := srv.userSrv.GetUserByEmail(email)
	if existing_user.UID != "" {
		return nil, errors.New("User already exist.")
	}

	country, _ := srv.userSrv.DefaultCountry()

	usr, _ := srv.userSrv.CreateUser(email, string(hashed_pass), country)
	token := srv.userSrv.GenerateToken(usr)

	srv.log.Info("New user token: ", token.Token)

	return usr, nil
}
