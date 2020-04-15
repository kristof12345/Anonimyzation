package anonmodel

import (
	"time"
)

// Kozponti tabla bejegyzes
type CentralTableItem struct {
	Id   int `bson:"id"`
	Time time.Time `bson:"time"`
}
