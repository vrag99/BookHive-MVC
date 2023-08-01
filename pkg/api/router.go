package api

import (
	"net/http"
	"github.com/gorilla/mux"
)

func Run(){
	r := mux.NewRouter()
    r.HandleFunc("/", hello).Methods("GET")
    http.ListenAndServe(":3000", r)
}

func hello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World!"))
}