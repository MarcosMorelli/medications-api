package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/MarcosMorelli/medication-api/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteMedication(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = database.MedicationCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		http.Error(w, "Failed to delete medication", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
