package handlers

import (
	"context"
	"encoding/json"

	"net/http"
	"time"

	"github.com/MarcosMorelli/medication-api/internal/database"
	"github.com/MarcosMorelli/medication-api/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateMedication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var medication models.Medication
	if err := json.NewDecoder(r.Body).Decode(&medication); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": medication}
	_, err = database.MedicationCollection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		http.Error(w, "Failed to update medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Medication updated successfully"}`))
}
