package web

import (
	"encoding/json"
	"github.com/sebastienfr/handsongo/dao"
	"github.com/sebastienfr/handsongo/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSpiritHandlerGet(t *testing.T) {

	// get mock dao
	daoMock, _ := dao.GetSpiritDAO("", dao.DAOMock)
	handler := NewSpiritHandler(daoMock)

	// build a request
	req, err := http.NewRequest("GET", "localhost/spirits", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	handler.GetAll(res, req)

	var spiritOut []model.Spirit
	json.NewDecoder(res.Body).Decode(&spiritOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	expected := model.Spirit{
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
	}

	if expected != spiritOut[0] {
		t.Errorf("Expected different from %v output %v", expected, spiritOut)
	}

	t.Logf("Test successfull expected \n %v \n is equal to result \n %v", expected, spiritOut[0])
}

func BenchmarkSpiritHandlerGet(t *testing.B) {

	// get mock dao
	daoMock, _ := dao.GetSpiritDAO("", dao.DAOMock)
	handler := NewSpiritHandler(daoMock)

	// build a request
	req, err := http.NewRequest("GET", "localhost/spirits", nil)
	if err != nil {
		t.Fatal(err)
	}

	// build the recorder
	res := httptest.NewRecorder()

	// execute the query
	handler.GetAll(res, req)

	var spiritOut []model.Spirit
	err = json.NewDecoder(res.Body).Decode(&spiritOut)

	if err != nil {
		t.Errorf("Unable to get JSON content %v", err)
	}

	expected := model.Spirit{
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
	}

	if expected != spiritOut[0] {
		t.Errorf("Expected different from %v output %v", expected, spiritOut)
	}
}
