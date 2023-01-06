package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitApi() {
	router := mux.NewRouter()
	router.HandleFunc("/api/infocards/{from}/{to}", GetInfoCardsByRange).Methods("GET")
	router.HandleFunc("/api/infocards/{type}", GetInfoCardsByType).Methods("GET")
	router.HandleFunc("/api/audio/key/{trackKey}", GetAudioByTrackKey).Methods("GET")
	router.HandleFunc("/api/audio/ref/{audioRef}", GetAudioByAudioRef).Methods("GET")
	http.ListenAndServe(":8000", router)
}
