package pickems

import (
	"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/common_http"
	mw "github.com/barqus/fillq_backend/internal/middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type HttpClient struct {
	svc Service
}

func MustNewHttpClient(service Service) *HttpClient {
	return &HttpClient{
		svc: service,
	}
}

func (c HttpClient) GetUsersPickems(w http.ResponseWriter, r *http.Request) {
	id, err := common_http.ParseURLParamToSInt(r, "user_id")

	usersPickems, err := c.svc.getUsersPickems(id)

	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, usersPickems)
}

func (c HttpClient) DeleteUsersAllPickems(w http.ResponseWriter, r *http.Request) {
	userId, err := common_http.ParseURLParamToSInt(r, "user_id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = isUserManipulatingHisPickems(userId, cookie.Value)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err = c.svc.deleteUsersAllPickems(userId)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}

func (c HttpClient) CreateUsersPickems(w http.ResponseWriter, r *http.Request) {
	currentDate := time.Now()
	location, err := time.LoadLocation("EET")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	closingDate := time.Date(2021, 12, 04, 12, 0, 0, 0, location)
	if currentDate.After(closingDate) == true {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var pickemData []*PickEm
	id, err := common_http.ParseURLParamToSInt(r, "user_id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	cookie, err := r.Cookie("jwt_token")
	logrus.Info("user_id", id)

	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = isUserManipulatingHisPickems(id, cookie.Value)
	if err != nil {

		common_http.WriteErrorResponse(w, err)
		return
	}

	if err = common_http.ParseJSON(r, &pickemData); err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err = c.svc.createUsersPickems(pickemData, id)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}

func isUserManipulatingHisPickems(userID int, jwtToken string) error {
	// TODO: TEST THIS OUT
	logrus.Info("jwtToken:", jwtToken)
	jwtTokenInformation, err := mw.VerifyToken(jwtToken)

	if err != nil {
		return err
	}

	jwtTokenMetadata := jwtTokenInformation.Claims.(jwt.MapClaims)
	authenticatedUserID := jwtTokenMetadata["user_id"].(string)

	intOfAuth, _ := strconv.Atoi(authenticatedUserID)
	logrus.Info("NOT YOUR", intOfAuth, " = ", userID)
	if intOfAuth != userID {
		return config.NOT_YOUR_PICKEMS
	}
	return nil
}
