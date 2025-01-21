package main

import (
	"log"

	"github.com/yosef-adel/toll-calculator/aggregator/client"
)

const (
	kafkatopic  = "obuData"
	aggEndpoint = "http://localhost:3000/aggregate"
)

func main() {
	calcSevice := NewClaculatorService()
	calcSevice = NewLogMiddleware(calcSevice)
	kafkaConsumer, err := NewKafkaConsumer(kafkatopic, calcSevice, client.NewClient(aggEndpoint))

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
