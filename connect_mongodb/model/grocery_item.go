package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroceryItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Quantity int                `bson:"quantity" json:"quantity"`
	Category string             `bson:"category" json:"category"`
}
