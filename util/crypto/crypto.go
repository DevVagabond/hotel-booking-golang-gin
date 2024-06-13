package crypto_util

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bytes, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		panic(error)
	}
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	error := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return error == nil
}
