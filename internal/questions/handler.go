package questions

import (
	"github.com/barqus/fillq_backend/internal/common_http"
	"net/http"
)

type HttpClient struct {
	svc Service
}

func MustNewHttpClient(service Service) *HttpClient {
	return &HttpClient{
		svc: service,
	}
}

func (c HttpClient) GetParticipantsQNA(w http.ResponseWriter, r *http.Request) {
	id, err := common_http.ParseURLParamToSInt(r, "participant_id")
	allParticipantsQNA, err := c.svc.getAllQnAByParticipantID(id)

	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, allParticipantsQNA)
}

func (c HttpClient) AddQNA(w http.ResponseWriter, r *http.Request) {
	var qnaData *QuestionAndAnswer

	if err := common_http.ParseJSON(r, &qnaData); err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err := c.svc.addNewQnAToParticipant(qnaData)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}

func (c HttpClient) DeleteQNAByID(w http.ResponseWriter, r *http.Request) {
	id, err := common_http.ParseURLParamToSInt(r, "id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err = c.svc.deleteQuestionByID(id)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}