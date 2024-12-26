package server

import (
	"net/http"

	"github.com/MarcosMorelli/medication-api/internal/handlers"
)

const V1_MEDICATIONS = "/v1/medications"

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HealthCheck)

	mux.Handle(V1_MEDICATIONS, MedicationsMux())

	return mux
}

func MedicationsMux() http.Handler {
	medicationsMux := http.NewServeMux()

	medicationsMux.HandleFunc("GET /", handlers.GetMedications)
	medicationsMux.HandleFunc("POST /", handlers.CreateMedication)

	medicationsMux.HandleFunc("GET /{id}", handlers.GetMedication)
	medicationsMux.HandleFunc("PUT /{id}", handlers.UpdateMedication)
	medicationsMux.HandleFunc("DELETE /{id}", handlers.DeleteMedication)

	return medicationsMux
}
