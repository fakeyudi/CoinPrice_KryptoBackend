package router

import (
	"CoinPrice_KryptoBackendTask/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()
	stdRouter := router.PathPrefix("/api").Subrouter()
	stdRouter.Use(middleware.JwtVerify)

	router.HandleFunc("/auth/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/auth/login", middleware.Login).Methods("POST", "OPTIONS")

	stdRouter.HandleFunc("/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
	stdRouter.HandleFunc("/users", middleware.GetAllUser).Methods("GET", "OPTIONS")
	stdRouter.HandleFunc("/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
	stdRouter.HandleFunc("/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

	stdRouter.HandleFunc("/createalert", middleware.CreateAlert).Methods("POST", "OPTIONS")
	stdRouter.HandleFunc("/getalert/{id}", middleware.GetAlert).Methods("GET", "OPTIONS")
	stdRouter.HandleFunc("/getuseralerts/{id}", middleware.GetAllUserAlerts).Methods("GET", "OPTIONS")
	stdRouter.HandleFunc("/deletealer/{id}", middleware.DeleteAlert).Methods("DELETE", "OPTIONS")
	stdRouter.HandleFunc("/updatealert/{id}", middleware.UpdateAlert).Methods("PUT","OPTIONS")

	return router
}
