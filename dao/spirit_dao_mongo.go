package dao

import (
	"github.com/sebastienfr/handsongo/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	collection = "spirits"
)

type SpiritDAOMongo struct {
	session *mgo.Session
}

func NewSpiritDAOMongo(session *mgo.Session) SpiritDAO {
	return &SpiritDAOMongo{
		session: session,
	}
}

func (s *SpiritDAOMongo) GetSpiritByID(ID string) (*model.Spirit, error) {
	spirit := model.Spirit{}
	c := s.session.DB("").C(collection)
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(ID)}).One(&spirit)
	return &spirit, err
}

func (s *SpiritDAOMongo) getAllSpiritsByQuery(query interface{}, start, end int) ([]model.Spirit, error) {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)

	// check param
	hasPaging := start > -1 && end > -1 && end > start

	// perform request
	var err error
	spirits := []model.Spirit{}
	if hasPaging {
		err = c.Find(query).Skip(start).Limit(end - start).All(&spirits)
	} else {
		err = c.Find(query).All(&spirits)
	}

	return spirits, err
}

func (s *SpiritDAOMongo) GetAllSpirits(start, end int) ([]model.Spirit, error) {
	return s.getAllSpiritsByQuery(nil, NoPaging, NoPaging)
}

func (s *SpiritDAOMongo) GetSpiritsByName(name string) ([]model.Spirit, error) {
	return s.getAllSpiritsByQuery(bson.M{"name": name}, NoPaging, NoPaging)
}

func (s *SpiritDAOMongo) GetSpiritsByType(spiritType string) ([]model.Spirit, error) {
	return s.getAllSpiritsByQuery(bson.M{"type": spiritType}, NoPaging, NoPaging)
}

func (s *SpiritDAOMongo) GetSpiritsByTypeAndScore(spiritType string, score uint8) ([]model.Spirit, error) {
	return s.getAllSpiritsByQuery(bson.M{"type": spiritType, "score": score}, NoPaging, NoPaging)
}

func (s *SpiritDAOMongo) SaveSpirit(spirit *model.Spirit) error {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	return c.Insert(spirit)
}

func (s *SpiritDAOMongo) UpsertSpirit(ID string, spirit *model.Spirit) (*mgo.ChangeInfo, error) {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	return c.Upsert(bson.M{"_id": bson.ObjectIdHex(ID)}, spirit)
}

func (s *SpiritDAOMongo) DeleteSpirit(ID string) error {
	session := s.session.Copy()
	defer session.Close()
	c := session.DB("").C(collection)
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(ID)})
	return err
}
