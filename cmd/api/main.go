package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/sivajithms/goapi/internal/handlers"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8000", r)

	if err != nil {
		log.Error(err)
	}
}
