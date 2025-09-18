package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashpassword), err

}

func Comparepassword(password, harshed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(harshed), []byte(password))
	return err == nil
}
