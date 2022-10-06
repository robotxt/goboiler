package core

import "golang.org/x/crypto/bcrypt"

func HashPassword(pwd []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}
