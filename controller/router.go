package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitApi() {
	router := mux.NewRouter()
	router.HandleFunc("/api/infocards/{from}/{to}", GetInfoCardsByRange).Methods("GET")
	router.HandleFunc("/api/infocards/{type}", GetInfoCardsByType).Methods("GET")
	router.HandleFunc("/api/audio/key/{trackKey}", GetAudioByTrackKey).Methods("GET")
	router.HandleFunc("/api/audio/ref/{audioRef}", GetAudioByAudioRef).Methods("GET")
	router.HandleFunc("/api/image/ref/{imageRef}", GetImageByImageRef).Methods("GET")
	router.HandleFunc("/api/infocards/{locale}", GetInfoCardsByLocale).Methods("GET")
	// router.HandleFunc("/api/infocards/{from}/{to}/{locale}", GetInfoCardsByLocaleAndYearRange).Methods("GET")
	http.ListenAndServe(":8000", router)
}
