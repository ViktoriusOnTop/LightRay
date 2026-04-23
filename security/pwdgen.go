package security

import (
	"github.com/sethvargo/go-password/password"
)

//This is to generate packages garble garble

func GenerateSecurePassword(passwordLength int) string {
	res, err := password.Generate(passwordLength, 0, 0, false, false)
	if err != nil {
		panic(err)
	}
	return res
}
