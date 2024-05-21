package exception

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	ps, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ps), nil
}

func ComparePassword(password []byte, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), password)
	return err == nil
}
