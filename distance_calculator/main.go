package main

import "log"

const kafkatopic = "obuData"

func main() {
	calcSevice := NewClaculatorService()
	kafkaConsumer, err := NewKafkaConsumer(kafkatopic, calcSevice)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
