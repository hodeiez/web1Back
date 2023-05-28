package controller

import (
	"bytes"
	"encoding/json"
	"hodei/web1/db"
	"hodei/web1/service"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetInfoCardsByRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	infos := service.GetInfoCardDTOByRange(params["from"], params["to"])
	json.NewEncoder(w).Encode(infos)
}
func GetInfoCardsByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	infos := service.GetInfoCardDTOByType(db.ToInfoType(params["type"]))
	json.NewEncoder(w).Encode(infos)
}
func GetAudioByTrackKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	audio := service.GetAudioFileByKey(params["trackKey"])
	http.ServeContent(w, r, "testAudio", time.Now(), bytes.NewReader(audio))
}
func GetAudioByAudioRef(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	audio := service.GetAudioFileByRef(params["audioRef"])
	http.ServeContent(w, r, params["audioRef"], time.Now(), bytes.NewReader(audio))
}
func GetImageByImageRef(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	image := service.GetImageFileByRef(params["imageRef"])
	http.ServeContent(w, r, params["imageRef"], time.Now(), bytes.NewReader(image))
}
func enableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		h.ServeHTTP(w, r)
	})
}

func GetInfoCardsByLocale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	infos := service.GetInfoCardDTOByLocale(params["locale"])
	json.NewEncoder(w).Encode(infos)
}
func GetInfoCardsByLocaleAndYearRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	infos := service.GetInfoCardDTOByLocaleAndYearRange(params["locale"], params["from"], params["to"])
	json.NewEncoder(w).Encode(infos)
}
