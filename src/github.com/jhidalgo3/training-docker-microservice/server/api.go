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
	r.HandleFunc("/api/info", getInfo).Methods("GET")

	//r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("./static/")))
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
}

func getConfig(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(config.Params)
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")

	p := config.Info{
		Instance: config.GetHostname(),
		Commit:   config.GetCommit(),
		Version:  config.GetVersion(),
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
