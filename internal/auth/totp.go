package auth 

import (
	"github.com/pquerna/otp/totp"
	"time"
)

func GenerateTOTP(secret, code string) bool {
	return totp.Validate(code, secret)
}

func GenerateSecret() (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Jagad Cipta Digital",
		AccountName: email,
	})
	if err != nil {
		return "", err
	}
	return key.Secret(), nil

}