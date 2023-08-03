package api

import (
	"net/http"
	"github.com/gorilla/mux"
	"BookHive/pkg/controller"
)

func Run(){
	r := mux.NewRouter()

	// Serving the static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

    r.HandleFunc("/", controller.HomePage).Methods("GET")

    http.ListenAndServe(":3000", r)
}