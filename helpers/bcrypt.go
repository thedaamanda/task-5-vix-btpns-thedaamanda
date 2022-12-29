package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePass := []byte(password)

	hashPass, err := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPass), nil
}

func ComparePassword(hashPass, password string) error {
	byteHashPass := []byte(hashPass)
	bytePass := []byte(password)

	return bcrypt.CompareHashAndPassword(byteHashPass, bytePass)
}
