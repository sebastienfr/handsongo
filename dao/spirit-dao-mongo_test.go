package dao

import (
	"github.com/sebastienfr/handsongo/model"
	"os"
	"testing"
	"time"
)

func TestDAOMongo(t *testing.T) {
	// get config
	config := os.Getenv("MONGODB_SRV")

	daoMongo, err := GetSpiritDAO(config, DAOMongo)
	if err != nil {
		t.Error(err)
	}

	toSave := model.Spirit{
		Name:         "Caroni 2000",
		Distiller:    "Caroni",
		Bottler:      "Velier",
		Country:      "Trinidad",
		Composition:  "Melasse",
		SpiritType:   model.TypeRhum,
		Age:          15,
		BottlingDate: time.Date(2015, 01, 01, 0, 0, 0, 0, time.UTC),
		Score:        8.5,
		Comment:      "heavy tire taste",
	}

	err = daoMongo.SaveSpirit(&toSave)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial spirit saved", toSave)

	spirits, err := daoMongo.GetAllSpirits(NoPaging, NoPaging)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial spirit found all", spirits[0])

	oneSpirit, err := daoMongo.GetSpiritByID(spirits[0].ID.Hex())
	if err != nil {
		t.Error(err)
	}

	t.Log("initial spirit found one", oneSpirit)

	oneSpirit.Age = 18
	oneSpirit.Comment = "soft tarmac smell"
	chg, err := daoMongo.UpsertSpirit(oneSpirit.ID.Hex(), oneSpirit)
	if err != nil {
		t.Error(err)
	}

	t.Log("initial spirit modified", chg, oneSpirit)

	oneSpirit, err = daoMongo.GetSpiritByID(oneSpirit.ID.Hex())
	if err != nil {
		t.Error(err)
	}

	t.Log("initial spirit found one modified", oneSpirit)

	err = daoMongo.DeleteSpirit(oneSpirit.ID.Hex())
	if err != nil {
		t.Error(err)
	}

	oneSpirit, err = daoMongo.GetSpiritByID(oneSpirit.ID.Hex())
	if err != nil {
		t.Log("initial spirit deleted", err, oneSpirit)
	}

}
