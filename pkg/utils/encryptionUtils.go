package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string{
	var saltRounds int = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func MatchPassword(inputPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}