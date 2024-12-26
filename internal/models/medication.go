package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Medication struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name   string             `bson:"name" json:"name"`
	Dosage string             `bson:"dosage" json:"dosage"`
	Form   string             `bson:"form" json:"form"`
}
