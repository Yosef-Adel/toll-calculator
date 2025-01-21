package main

import (
	"github.com/yosef-adel/toll-calculator/types"
)

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAgregator struct {
	store Storer
}

func NewInvoiceAgregator(store Storer) Aggregator {
	return &InvoiceAgregator{
		store: store,
	}
}

func (i *InvoiceAgregator) AggregateDistance(dist types.Distance) error {
	return i.store.Insert(dist)
}
