package dao

import (
	"github.com/sebastienfr/handsongo/model"
	"gopkg.in/mgo.v2"
)

const (
	// NoPaging used with skip, limit parameters
	NoPaging = -1
)

// SpiritDAO is the DAO interface to work with spirits
type SpiritDAO interface {
	//
	GetSpiritByID(ID string) (*model.Spirit, error)
	//
	GetAllSpirits(skip, limit int) ([]model.Spirit, error)
	//
	GetSpiritsByName(name string) ([]model.Spirit, error)
	//
	GetSpiritsByType(spiritType string) ([]model.Spirit, error)
	//
	GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error)
	//
	SaveSpirit(spirit *model.Spirit) error
	//
	UpsertSpirit(ID string, spirit *model.Spirit) (*mgo.ChangeInfo, error)
	//
	DeleteSpirit(ID string) error
}
