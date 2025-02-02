package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yosef-adel/toll-calculator/types"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CalculateDistance(data types.OBUData) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": dist,
		}).Info("Calculate distance")
	}(time.Now())

	return l.next.CalculateDistance(data)
}
