package anonmodel

import "fmt"

// Eklivalencia osztaly
type EqulivalenceClass struct {
	Id                  int
	CategoricAttributes map[string]string
	IntervalAttributes  map[string]NumericRange
	Count               int
	Active              bool
}

func (e EqulivalenceClass) Print() {
	fmt.Println(e.Id)

	for _, attr := range e.IntervalAttributes {
		fmt.Println(attr)
	}

	for _, attr := range e.CategoricAttributes {
		fmt.Println(attr)
	}
}
