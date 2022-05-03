package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/ahmedmahmo/learn/booking/users"
	Mux "github.com/gorilla/mux"
	gohandlers "github.com/gorilla/handlers"
)

const APP = "booking-api "
func main() {
	// Init logging
	l := log.New(os.Stdout, APP, log.LstdFlags)
	
	// HTTP Handlers
	usersHandler    := serve.NewUsers(l)

	// Router
	mx := Mux.NewRouter()

	// Handling all GET Requests
	get := mx.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {l.Println(http.MethodGet)})
	get.HandleFunc("/users", usersHandler.Get)


	// Handling all POST Requests
	post := mx.Methods(http.MethodPost).Subrouter()
	post.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	post.HandleFunc("/users", usersHandler.Post)

	post.Use(usersHandler.ValidateUsersMiddleware)
	// post.Use(accountsHandler.ValidateUsersMiddleware)

	// Handling all PUT Requests
	put := mx.Methods(http.MethodPut).Subrouter()
	put.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	put.HandleFunc("/users/{id:[0-9]+}", usersHandler.Put)

	put.Use(usersHandler.ValidateUsersMiddleware)
	// put.Use(accountsHandler.ValidateUsersMiddleware)

	// Handling all DELETE Requests
	delete := mx.Methods(http.MethodDelete).Subrouter()
	delete.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	delete.HandleFunc("/users/{id:[0-9]+}", usersHandler.Delete)

	// Handle CORS
	gh := gohandlers.CORS(gohandlers.AllowedOrigins([]string{
		"*",
	}))

	// Init Server with Timeouts specs
	server := &http.Server{
		Addr: ":8080",
		Handler: gh(mx),
		IdleTimeout: 120 *time.Second,
		ReadTimeout: 1   *time.Second,
		WriteTimeout: 1  *time.Second,
	}
	go func() {
		l.Println("Starting server on :8080")
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved Terminate, graceful shutdown", sig)
	tc, cancel:= context.WithTimeout(context.Background(), 30 *time.Second)
	defer cancel()
	server.Shutdown(tc)
}