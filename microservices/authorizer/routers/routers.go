package routers

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	mainRouter := mux.NewRouter()
	authRouter := mainRouter.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", authorizer.SignupHandler).Methods("POST")
	authRouter.HandleFunc("/signin", authorizer.SigninHandler).Methods("GET")
	return mainRouter
}
