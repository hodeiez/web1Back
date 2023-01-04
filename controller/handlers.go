package controller

import (
	"encoding/json"
	"hodei/web1/db"
	"hodei/web1/service"
	"net/http"

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
