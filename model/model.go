package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	//
	TypeRhum = "rhum"
	//
	TypeWhine = "wine"
	//
	TypeBeer = "beer"
	//
	TypeCalados = "calvados"
	//
	TypeChampagne = "champagne"
	//
	TypeGin = "gin"
)

// Spirit is the struture to define a spirit
type Spirit struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty" `
	Name         string        `json:"name" bson:"name"`
	Distiller    string        `json:"distiller" bson:"distiller"`
	Bottler      string        `json:"bottler" bson:"bottler"`
	Country      string        `json:"country" bson:"country"`
	Region       string        `json:"region" bson:"region"`
	Composition  string        `json:"composition" bson:"composition"`
	SpiritType   string        `json:"type" bson:"type"`
	Age          uint8         `json:"distillationYear" bson:"distillationYear"`
	BottlingDate time.Time     `json:"bottlingDate" bson:"bottlingDate"`
	Score        float32       `json:"score" bson:"score"`
	Comment      string        `json:"comment" bson:"comment"`
}
