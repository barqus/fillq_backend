package api

import (
	"encoding/json"
	"net/http"
)

type HttpError struct {
	Message string `json:"error"`
}

type HttpSuccess struct {
	Data interface{} `json:"data"`
}

func (e HttpSuccess) Encode(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(e); err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("internal server error"))
	}
}

func (e HttpError) Encode(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(e); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}