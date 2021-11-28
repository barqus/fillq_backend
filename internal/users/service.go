package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/common_http"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type Service interface {
	loginUser(code string) (*User, error)
	getUserByID(id string) (*User, error)
}

type service struct {
	clt     common_http.Client
	storage Storage
}

func MustNewService(client common_http.Client, storage Storage) *service {
	return &service{
		client,
		storage,
	}
}

func (s *service) getUserByID(id string) (*User, error) {
	foundUser, err := s.storage.getUserByID(id)
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
func (s *service) loginUser(code string) (*User, error) {

	existingUser, err := s.storage.getUserByCode(code)
	if err != nil {
		if err != config.USER_NOT_FOUND {
			logrus.Error("FAILED TO GET USER BY CODE " + code)
			return nil, err
		}
	}
	if existingUser != nil {
		return existingUser, nil
	}

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	grantType := os.Getenv("GRANT_TYPE")
	redirectURI := os.Getenv("REDIRECT_URI")
	twitchLoginUrl := fmt.Sprintf("https://id.twitch.tv/oauth2/token?client_id=%s&client_secret=%s&code=%s&grant_type=%s&redirect_uri=%s", clientID, clientSecret, code, grantType, redirectURI)

	ctx := context.TODO()
	rawResponse, statusCode, err := s.clt.Post(ctx, twitchLoginUrl, "", nil, nil)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, config.TWITCH_CODE_INVALID
	}
	var responseObject UserAccessInformation

	err = json.Unmarshal(rawResponse, &responseObject)
	if err != nil {
		return nil, err
	}

	userTwitchURI := "https://api.twitch.tv/helix/users"
	rawResponse, statusCode, err = s.clt.Get(ctx, userTwitchURI, responseObject.AccessToken, clientID, nil)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, config.TWITCH_CODE_INVALID
	}

	var responseTwitchUserInfoObject TwitchData
	err = json.Unmarshal(rawResponse, &responseTwitchUserInfoObject)
	if err != nil {
		return nil, err
	}

	var twitchChannelInfo TwitchChannelInformation
	twitchChannelInfo = responseTwitchUserInfoObject.Data[0]

	jwtToken, err := generateJWTForUser(twitchChannelInfo.ID)
	if err != nil {
		return nil, err
	}

	userInformation := &User{
		ID:              twitchChannelInfo.ID,
		DisplayName:     twitchChannelInfo.DisplayName,
		ProfileImageURL: twitchChannelInfo.ProfileImageURL,
		Email:           twitchChannelInfo.Email,
		TwitchCode:      code,
		Role:            "regular",
		AccessToken: responseObject.AccessToken,
		RefreshToken: responseObject.RefreshToken,
		JWTToken: jwtToken,
	}

	existingUser, err = s.storage.getUserByID(userInformation.ID)
	if err != nil {
		if err != config.USER_NOT_FOUND {
			logrus.Error("FAILED TO GET USER BY ID " + userInformation.ID)
			return nil, err
		}
	}
	if existingUser != nil {
		logrus.Info("USER EXISTS")
		return existingUser, nil
	}

	logrus.Info("ADDING USER")
	err = s.storage.addNewUser(userInformation)
	if err != nil {
		logrus.Error("FAILED TO ADD USER" + userInformation.DisplayName)
		return nil, err
	}
	logrus.Info("ADDED USER: ", userInformation.DisplayName)
	return userInformation, nil
}

func generateJWTForUser(userID string) (string, error){
	var err error
	atClaims := jwt.MapClaims{}
	if userID == "107170813" {
		atClaims["admin"] = true
	} else {
		atClaims["admin"] = false
	}

	atClaims["user_id"] = userID
	expirationDate := time.Now().AddDate(0,1,0)
	fmt.Println(expirationDate)
	atClaims["exp"] = expirationDate
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	logrus.Info("HERER",err,token)
	if err != nil {
		return "", err
	}
	return token, nil
}