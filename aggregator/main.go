package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yosef-adel/toll-calculator/types"
)

func main() {
	store := NewMemorySotre()
	svc := NewInvoiceAgregator(store)
	svc = NewLogMiddleWare(svc)
	makeHTTPTransport(svc)
}

func makeHTTPTransport(svc Aggregator) {
	fmt.Println("HTTP transport running in port 3000")

	http.HandleFunc("/aggregate", handleAggregate(svc))
	_ = http.ListenAndServe(":3000", nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			_ = writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			_ = writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
