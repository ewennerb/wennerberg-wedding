package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Guest struct {
	Name               string `bson:"name" json:"name"`
	Attending          *bool  `bson:"attending" json:"attending"`
	DinnerChoice       string `bson:"dinner_choice" json:"dinner_choice"`
	DietaryRestriction string `bson:"dietary_restriction" json:"dietary_restriction"`
}

type Household struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Names  []string           `bson:"names" json:"names"`
	Guests []Guest            `bson:"guests" json:"guests"`
}
