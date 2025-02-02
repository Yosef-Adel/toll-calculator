obu:
	@go build -o bin/obu obu/main.go
	@./bin/obu

receiver:
	@go build -o bin/receiver ./data_receiver
	@./bin/receiver

calc:
	@go build -o bin/calc ./distance_calculator
	@./bin/calc

agg:
	@go build -o bin/aggregator ./aggregator
	@./bin/aggregator

.PHONY: obu invoicer

