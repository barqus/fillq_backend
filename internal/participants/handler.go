package participants

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

func (c HttpClient) GetAllParticipants(w http.ResponseWriter, r *http.Request) {
	allParticipants, err := c.svc.getAllParticipants()

	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, allParticipants)
}

func (c HttpClient) AddParticipant(w http.ResponseWriter, r *http.Request) {
	var participantData *Participant

	if err := common_http.ParseJSON(r, &participantData); err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err := c.svc.addParticipants(participantData)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}

func (c HttpClient) DeleteParticipant(w http.ResponseWriter, r *http.Request) {
	id, err := common_http.ParseURLParamToSInt(r, "id")
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	err = c.svc.deleteParticipant(id)
	if err != nil {
		common_http.WriteErrorResponse(w, err)
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, 200)
}


//func AddParticipant(w http.ResponseWriter, r *http.Request) {
//	// Read to request body
//	defer r.Body.Close()
//	body, err := ioutil.ReadAll(r.Body)
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	var participant Participant
//	json.Unmarshal(body, &participant)
//
//	// Append to the Book mocks
//	participant.Id = rand.Intn(100)
//	ParticipantsMock = append(ParticipantsMock, participant)
//
//	// Send a 201 created response
//	w.WriteHeader(http.StatusCreated)
//	w.Header().Add("Content-Type", "application/json")
//	json.NewEncoder(w).Encode("Created")
//}
