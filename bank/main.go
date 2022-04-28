package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/ahmedmahmo/learn/bank/handlers"
	"github.com/gorilla/mux"
)

const timeout = 30 * time.Second

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	
	ph := handlers.NewProduct(l)

	sm := mux.NewRouter()
	sm.Handle("/products", ph)

	server := &http.Server{
		Addr: ":8080",
		Handler: sm,
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
	tc, cancel:= context.WithTimeout(context.Background(), timeout)
	defer cancel()
	server.Shutdown(tc)
}