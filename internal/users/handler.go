package users

import (
	"github.com/barqus/fillq_backend/internal/common_http"
	"net/http"
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

func (c HttpClient) LoginUser(w http.ResponseWriter, r *http.Request) {
	code, err := common_http.ParseURLParamToString(r, "id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	loggedUser, err := c.svc.loginUser(code)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	expireDate := time.Now().AddDate(0,0,30)
	http.SetCookie(w, &http.Cookie{
		Name: "jwt_token",
		Value: loggedUser.JWTToken,
		Path: "/",
		Expires: expireDate,
	})
	common_http.WriteJSONResponse(w, http.StatusOK, &loggedUser.ID)
}


func (c HttpClient) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := common_http.ParseURLParamToString(r, "id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	FoundUser, err := c.svc.getUserByID(id)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, FoundUser)
}
