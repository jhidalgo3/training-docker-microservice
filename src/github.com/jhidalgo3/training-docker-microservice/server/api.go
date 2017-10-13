package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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

	hostname := getHostname()
	p := config.Info{
		Instance: hostname,
		Version:  getVersion(),
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getVersion() string {
	ver := config.Version
	if ver == "" {
		ver = "-"
	}

	return ver
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return hostname
}
