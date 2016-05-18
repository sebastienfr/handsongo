package dao

import (
	"github.com/sebastienfr/handsongo/model"
	"gopkg.in/mgo.v2"
	"time"
)

// SpiritDAOMock is the mocked implementation of the SpiritDAO
type SpiritDAOMock struct {
}

// NewSpiritDAOMock creates a new SpiritDAO with a mocked implementation
func NewSpiritDAOMock() SpiritDAO {
	return &SpiritDAOMock{}
}

// GetSpiritByID returns a spirit by its ID
func (s *SpiritDAOMock) GetSpiritByID(ID string) (*model.Spirit, error) {
	return nil, nil
}

// GetAllSpirits returns all spirits with paging capability
func (s *SpiritDAOMock) GetAllSpirits(start, end int) ([]model.Spirit, error) {
	return []model.Spirit{
		model.Spirit{
			Name:         "Caroni",
			Distiller:    "Caroni",
			Bottler:      "Velier",
			Country:      "Trinidad",
			Composition:  "Molasse",
			SpiritType:   model.TypeRhum,
			Age:          15,
			BottlingDate: time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
			Score:        8.5,
			Comment:      "heavy tire taste",
		},
	}, nil
}

// GetSpiritsByName returns all spirits by name
func (s *SpiritDAOMock) GetSpiritsByName(name string) ([]model.Spirit, error) {
	return nil, nil
}

// GetSpiritsByType returns all spirits by type
func (s *SpiritDAOMock) GetSpiritsByType(spiritType string) ([]model.Spirit, error) {
	return nil, nil
}

// GetSpiritsByTypeAndScore returns all spirits by type and score greater than parameter
func (s *SpiritDAOMock) GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error) {
	return nil, nil
}

// SaveSpirit saves the spirit
func (s *SpiritDAOMock) SaveSpirit(spirit *model.Spirit) error {
	return nil
}

// UpsertSpirit updates or creates a spirit
func (s *SpiritDAOMock) UpsertSpirit(ID string, spirit *model.Spirit) (*mgo.ChangeInfo, error) {
	return nil, nil
}

// DeleteSpirit deletes a spirits by its ID
func (s *SpiritDAOMock) DeleteSpirit(ID string) error {
	return nil
}
