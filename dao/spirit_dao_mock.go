package dao

import "github.com/sebastienfr/handsongo/model"

type SpiritDAOMock struct {
}

func NewSpiritDAOMock() SpiritDAO {
	return &SpiritDAOMongo{}
}

func (s *SpiritDAOMock) GetSpiritByID(ID string) (*model.Spirit, error) {
	return nil, nil
}

func (s *SpiritDAOMock) GetAllSpirits(skip, limit int) ([]model.Spirit, error) {
	return nil, nil
}

func (s *SpiritDAOMock) GetSpiritsByName(name string) ([]model.Spirit, error) {
	return nil, nil
}

func (s *SpiritDAOMock) GetSpiritsByType(spiritType string) ([]model.Spirit, error) {
	return nil, nil
}

func (s *SpiritDAOMock) GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error) {
	return nil, nil
}

func (s *SpiritDAOMock) SaveSpirit(spirit *model.Spirit) error {
	return nil
}

func (s *SpiritDAOMock) UpsertSpirit(ID string, spirit *model.Spirit) error {
	return nil
}

func (s *SpiritDAOMock) DeleteSpirit(ID string) error {
	return nil
}
