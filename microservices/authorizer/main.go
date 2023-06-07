package main

import (
	"fmt"
	"net/http"

	"authorizer/routers"
	gohandlers "github.com/gorilla/handlers"
)

func main() {
	router := routers.RegisterRoutes()
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:3000"}))
	// Add the Middleware to different subrouter
	// HTTP Server
	// Add Time outs
	server := &http.Server{
		Addr:    "127.0.0.1:9090",
		Handler: ch(router),
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error Booting the Server")
	}
}
