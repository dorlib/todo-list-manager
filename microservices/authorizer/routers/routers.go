package routers

import (
	"authorizer/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	authRouter.HandleFunc("/signin", handlers.SigninHandler).Methods("GET")
	authRouter.HandleFunc("/validate-url", handlers.ValidateHandler).Methods("GET")
	return mainRouter
}
