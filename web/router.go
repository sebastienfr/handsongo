package web

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/sebastienfr/handsongo/utils"
	"net/http"
)

// Router is the struct use for routing
type Router struct {
	*mux.Router
}

// Route is a structure of Route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// NewRouter creates a new router instance
func NewRouter(handler *SpiritHandler) *Router {
	// new router
	router := Router{mux.NewRouter()}

	// default JSON not found handler
	router.NotFoundHandler = utils.NotFoundHandler()

	// no strict slash
	router.StrictSlash(false)

	// add routes of handler
	for _, route := range handler.Routes {
		logger.WithField("route", route).Debug("adding route to mux")
		router.
			Methods(route.Method).
			Path(handler.Prefix + route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return &router
}
