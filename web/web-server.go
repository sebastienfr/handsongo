package web

import (
	logger "github.com/Sirupsen/logrus"
	"github.com/meatballhat/negroni-logrus"
	"github.com/sebastienfr/handsongo/dao"
	"github.com/urfave/negroni"
	"time"
)

// BuildWebServer constructs a new web server with the right DAO and spirits handler
func BuildWebServer(db string, daoType int, statisticsDuration time.Duration) (*negroni.Negroni, error) {

	// spirit dao
	dao, err := dao.GetSpiritDAO(db, daoType)
	if err != nil {
		logger.WithField("error", err).WithField("connection string", db).Fatal("unable to connect to mongo db")
		return nil, err
	}

	// web server
	n := negroni.New()

	// new handler
	handler := NewSpiritHandler(dao)

	// new router
	router := NewRouter(handler)

	// add middleware for logging
	n.Use(negronilogrus.NewMiddlewareFromLogger(logger.StandardLogger(), "spirit"))

	// add recovery middleware in case of panic in handler func
	recovery := negroni.NewRecovery()
	recovery.PrintStack = false
	n.Use(recovery)

	// add statistics middleware
	n.Use(NewStatisticsMiddleware(statisticsDuration))

	// add as many middleware as you like

	// handler goes last
	n.UseHandler(router)

	return n, nil
}
