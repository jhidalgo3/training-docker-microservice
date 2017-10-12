package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhidalgo3/training-docker-microservice/config"
	"github.com/spf13/viper"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/api/config", getConfig).Methods("GET")

	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(config.Params)
}
