package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"os"
)

func VerifyToken(jwtToken string) (*jwt.Token, error) {
	logrus.Info("Parsing JWT TOKEN:", jwtToken)
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logrus.Error("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return token, nil
}
