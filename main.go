package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const serverPort = "8000"

func main() {
	// initialize new router
	r := chi.NewRouter()

	// middlewares
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Welcome"))
	})

	// define server and start it

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: r,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Print("The authentication server could not start !")
		log.Panic(err)
	}
}
