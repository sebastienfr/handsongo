package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	// TypeRhum is the constant used for rhum spirits
	TypeRhum = "rhum"
	// TypeWhine is the constant used for rhum spirits
	TypeWhine = "wine"
	// TypeBeer is the constant used for rhum spirits
	TypeBeer = "beer"
	// TypeCalados is the constant used for rhum spirits
	TypeCalados = "calvados"
	// TypeChampagne is the constant used for rhum spirits
	TypeChampagne = "champagne"
	// TypeGin is the constant used for rhum spirits
	TypeGin = "gin"
)

// Spirit is the structure to define a spirit
type Spirit struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Name         string        `json:"name" bson:"name"`
	Distiller    string        `json:"distiller" bson:"distiller"`
	Bottler      string        `json:"bottler" bson:"bottler"`
	Country      string        `json:"country" bson:"country"`
	Region       string        `json:"region" bson:"region"`
	Composition  string        `json:"composition" bson:"composition"`
	SpiritType   string        `json:"type" bson:"type"`
	Age          uint8         `json:"age" bson:"age"`
	BottlingDate time.Time     `json:"bottlingDate" bson:"bottlingDate"`
	Score        float32       `json:"score" bson:"score"`
	Comment      string        `json:"comment" bson:"comment"`
}

// GetID returns the ID of an Spirit as a string
func (s *Spirit) GetID() string {
	return s.ID.Hex()
}
