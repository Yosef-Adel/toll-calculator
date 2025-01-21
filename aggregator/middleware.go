package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yosef-adel/toll-calculator/types"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleWare(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) AggregateDistance(dist types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("Agregate Distance")

	}(time.Now())

	return l.next.AggregateDistance(dist)
}
