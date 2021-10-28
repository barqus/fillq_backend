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
		return
	}

	common_http.WriteJSONResponse(w, http.StatusOK, allParticipants)
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(ParticipantsMock)
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
