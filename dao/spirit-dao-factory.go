package dao

import (
	"errors"
	"gopkg.in/mgo.v2"
	"time"
)

const (
	// DAOMongo is used for Mongo implementation of SpiritDAO
	DAOMongo int = iota
	// DAOMock is used for mocked implementation of SpiritDAO
	DAOMock

	// mongo timeout
	timeout = 15 * time.Second
	// poolSize of mongo connection pool
	poolSize = 35
)

var (
	// ErrorDAONotFound is used for unknown DAO type
	ErrorDAONotFound = errors.New("Unknown DAO type")
)

// GetSpiritDAO returns a SpiritDAO according to type and params
func GetSpiritDAO(param string, daoType int) (SpiritDAO, error) {
	switch daoType {
	case DAOMongo:
		// mongo connection
		mgoSession, err := mgo.DialWithTimeout(param, timeout)
		if err != nil {
			return nil, err
		}

		// set 30 sec timeout on session
		mgoSession.SetSyncTimeout(timeout)
		mgoSession.SetSocketTimeout(timeout)
		// set mode
		mgoSession.SetMode(mgo.Monotonic, true)
		mgoSession.SetPoolLimit(poolSize)

		return NewSpiritDAOMongo(mgoSession), nil
	case DAOMock:
		return NewSpiritDAOMock(), nil
	default:
		return nil, ErrorDAONotFound
	}
}
