package web

import (
	"github.com/sebastienfr/handsongo/utils"
	"net/http"
	"time"
)

type StatisticsMiddleware struct {
	Stat *utils.Statistics
}

func NewStatisticsMiddleware(duration time.Duration) *StatisticsMiddleware {
	return &StatisticsMiddleware{
		Stat: utils.NewStatistics(duration),
	}
}

func (sm *StatisticsMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	sm.Stat.PlusOne()
	next(rw, r)
}
