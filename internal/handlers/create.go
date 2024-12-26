package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/MarcosMorelli/medication-api/internal/database"
	"github.com/MarcosMorelli/medication-api/internal/models"
)

func CreateMedication(w http.ResponseWriter, r *http.Request) {
	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.MedicationCollection.InsertOne(ctx, medication)
	if err != nil {
		http.Error(w, "Failed to create medication", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{"id": result.InsertedID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
