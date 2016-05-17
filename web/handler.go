package web

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/sebastienfr/handsongo/dao"
	"github.com/sebastienfr/handsongo/utils"
	"net/http"
)

const (
	prefix = "/spirits"
)

// SpiritHandler is a boxes handler
type SpiritHandler struct {
	spiritDao dao.SpiritDAO
	Routes    []Route
	Prefix    string
}

func NewSpiritHandler(spriritDAO dao.SpiritDAO) *SpiritHandler {
	handler := SpiritHandler{
		spiritDao: spriritDAO,
		Prefix:    prefix,
	}

	// build routes
	routes := []Route{}
	// GetAll
	routes = append(routes, Route{
		Name:        "Get all spirits",
		Method:      "GET",
		Pattern:     "",
		HandlerFunc: handler.GetAll,
	})
	// Get
	routes = append(routes, Route{
		Name:        "Get one spirit",
		Method:      "GET",
		Pattern:     "/{id}",
		HandlerFunc: handler.Get,
	})
	// Create
	routes = append(routes, Route{
		Name:        "Create a spirit",
		Method:      "POST",
		Pattern:     "",
		HandlerFunc: handler.Create,
	})
	// Update
	routes = append(routes, Route{
		Name:        "Update a spirit",
		Method:      "PUT",
		Pattern:     "/{id}",
		HandlerFunc: handler.Update,
	})
	// Delete
	routes = append(routes, Route{
		Name:        "Delete a spirit",
		Method:      "DELETE",
		Pattern:     "/{id}",
		HandlerFunc: handler.Delete,
	})

	handler.Routes = routes

	return &handler
}

// GetAll retrieve all entities with optional paging of items (start / end are item counts 50 to 100 for example)
func (sh *SpiritHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO process range params

	// find all spirits
	spirits, err := sh.spiritDao.GetAllSpirits(dao.NoPaging, dao.NoPaging)
	if err != nil {
		logger.WithField("error", err).Warn("unable to retrieve spirits")
		utils.SendJSONError(w, "Error while retrieving spirits", http.StatusInternalServerError)
		return
	}

	logger.WithField("spirits", spirits).Debug("spirits found")
	utils.SendJSONOk(w, spirits)
}

// Get retrieve an entity by id
func (sh *SpiritHandler) Get(w http.ResponseWriter, r *http.Request) {
	// get the spirit ID from the URL
	spiritID := utils.ParamAsString("id", r)

	// find spirit
	spirit, err := sh.spiritDao.GetSpiritByID(spiritID)
	if err != nil {
		logger.WithField("error", err).WithField("spirit ID", spiritID).Warn("unable to retrieve spirit by ID")
		utils.SendJSONError(w, "Error while retrieving spirit by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("spirits", spirit).Debug("spirit found")
	utils.SendJSONOk(w, spirit)
}

// Create create an entity
func (sh *SpiritHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
	utils.SendJSONError(w, "Not Implemented", http.StatusNotImplemented)
}

// Update update an entity by id
func (sh *SpiritHandler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO
	utils.SendJSONError(w, "Not Implemented", http.StatusNotImplemented)
}

// Delete delete an entity by id
func (sh *SpiritHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO
	utils.SendJSONError(w, "Not Implemented", http.StatusNotImplemented)
}
