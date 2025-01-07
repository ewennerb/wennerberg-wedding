package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	Name         string `bson:"name" json:"name"`
	Attending    string `bson:"attending" json:"attending"`
	DinnerChoice string `bson:"dinner_choice" json:"dinner_choice"`
}

type Household struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Names  []string           `bson:"names" json:"names"`
	Guests []Guest            `bson:"guests" json:"guests"`
	Notes  string             `bson:"notes" json:"notes"`
}
