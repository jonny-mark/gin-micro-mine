/**
 * @author jiangshangfang
 * @date 2021/12/8 10:15 PM
 **/
package auth

import "golang.org/x/crypto/bcrypt"

func HashFromPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CompareHashAndPassword(hashed string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}
