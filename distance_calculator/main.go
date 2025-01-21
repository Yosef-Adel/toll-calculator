package main

import "log"

const kafkatopic = "obuData"

func main() {
	calcSevice := NewClaculatorService()
	calcSevice = NewLogMiddleware(calcSevice)
	kafkaConsumer, err := NewKafkaConsumer(kafkatopic, calcSevice)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
