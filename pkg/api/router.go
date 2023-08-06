package api

import (
	"BookHive/pkg/controller"
	"BookHive/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func Run(){
	r := mux.NewRouter()

	// Serving the static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/static/").Handler(s)

	r.Use(utils.ValidateJWT)

	// Handling home page
    r.HandleFunc("/", controller.HomePage).Methods("GET")

	// Handling register page
    r.HandleFunc("/register", controller.RegisterPage).Methods("GET")
    r.HandleFunc("/register", controller.SignUpRequest).Methods("POST")

	// Handling login page
	r.HandleFunc("/login", controller.LoginPage).Methods("GET")
	r.HandleFunc("/login", controller.LoginRequest).Methods("POST")

    http.ListenAndServe(":3000", r)
}