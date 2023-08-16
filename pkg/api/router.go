package api

import (
	"BookHive/pkg/controller"
	"BookHive/pkg/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	// Serving the static files
	s := http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))
	router.PathPrefix("/static/").Handler(s)

	// Middleware for validating jwt
	router.Use(middleware.ValidateJWT)

	// Handling home page
	router.HandleFunc("/", controller.HomePage).Methods("GET")

	// Handling register page
	router.HandleFunc("/register", controller.RegisterPage).Methods("GET")
	router.HandleFunc("/register", controller.SignUpRequest).Methods("POST")

	// Handling login page
	router.HandleFunc("/login", controller.LoginPage).Methods("GET")
	router.HandleFunc("/login", controller.LoginRequest).Methods("POST")

	adminRouter := router.PathPrefix("/adminDashboard").Subrouter()
	userRouter := router.PathPrefix("/userDashboard").Subrouter()

	adminRouter.Use(middleware.CheckForAdmin(true))
	userRouter.Use(middleware.CheckForAdmin(false))

	// Handling UserDashboard
	userRouter.HandleFunc("", controller.UserViews).Methods("GET")
	userRouter.HandleFunc("/{viewMode}", controller.UserViews).Methods("GET")
	userRouter.HandleFunc("/request/{id}", controller.RequestBook).Methods("GET")
	userRouter.HandleFunc("/requestReturn/{id}", controller.RequestReturnBook).Methods("GET")

	// Handling AdminDashboard
	adminRouter.HandleFunc("", controller.AdminViews).Methods("GET")
	adminRouter.HandleFunc("/deleteBook/{id}", controller.DeleteBook).Methods("GET")
	adminRouter.HandleFunc("/addBook", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/issueRequests", controller.IssueRequests).Methods("GET")
	adminRouter.HandleFunc("/issueRequests/{action}/{id}", controller.IssueRequests).Methods("GET")
	adminRouter.HandleFunc("/returnRequests", controller.ReturnRequests).Methods("GET")
	adminRouter.HandleFunc("/returnRequests/{action}/{id}", controller.ReturnRequests).Methods("GET")
	adminRouter.HandleFunc("/adminRequests", controller.AdminRequests).Methods("GET")
	adminRouter.HandleFunc("/adminRequests/{action}/{id}", controller.AdminRequests).Methods("GET")

	//Logout
	router.HandleFunc("/logout", controller.Logout).Methods("GET")

	//404
	router.NotFoundHandler = http.HandlerFunc(controller.NotFound)

	//403
	router.HandleFunc("/forbiddenRequest", controller.ForbiddenRequest).Methods("GET")

	//500
	router.HandleFunc("/internalServerError", controller.InternalServerError).Methods("GET")

	fmt.Println("Listening on http://localhost:3000")
	http.ListenAndServe(":3000", router)
}
