package auth

import "golang.org/x/crypto/bcrypt"

// HashFromPassword 加密
func HashFromPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// CompareHashAndPassword 一致性校验
func CompareHashAndPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
