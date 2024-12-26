package handlers

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"

	"net/http"
	"time"

	"github.com/MarcosMorelli/medication-api/internal/database"
	"github.com/MarcosMorelli/medication-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMedications(w http.ResponseWriter, r *http.Request) {
	slog.Debug("GetMedications")
	var medications []models.Medication

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.MedicationCollection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Println("Error finding medications:", err)
		http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var medication models.Medication
		if err := cursor.Decode(&medication); err != nil {
			log.Println("Error decoding medication:", err)
			http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
			return
		}
		medications = append(medications, medication)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		http.Error(w, "Failed to fetch medications", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medications)
}

func GetMedication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var medication models.Medication

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = database.MedicationCollection.FindOne(ctx, bson.M{"_id": objID}).Decode(&medication)
	if err != nil {
		http.Error(w, "Medication not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(medication)
}
