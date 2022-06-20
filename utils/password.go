package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(passwrod string) (string,error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(passwrod),14)
	return string(hashpass),err
}

func ComparePassword(hashpass string,password string) bool  {

	err := bcrypt.CompareHashAndPassword([]byte(hashpass),[]byte(password))
	return err==nil
	
}