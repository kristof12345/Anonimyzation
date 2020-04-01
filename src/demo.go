package main

import (
	"anonbll"
	"anondb"
	"anonmodel"
	"fmt"
)

// Dataset: lenyegeben egy adatbazis tabla,
// amire anonimizalasi beallitasokat adunk meg (k, algoritmus, folytonos-e)
// es mezoi vannak (fields)

//Field egy adatbazis oszlop,
//van neve, modja ('id', 'qid', 'keep' or 'drop') es tipusa ('numeric' or 'prefix' or 'coords')

//A routers fajlban vannak az elerheto url-ek es a funkciokat megvalosito fuggvenyek nevei.

// A document reprezental egy konkret adatot / rekordot
// kulcs - ertek parokban tarolja a mezo nevet es erteket

func main() {
	fmt.Printf("APP STARTED\n")

	// Document
	var document = make(map[string]interface{})
	document["City"] = "Bp"
	document["Age"] = anonmodel.NumericRange{40, 50}

	var eqs = anonbll.GetMatchingClasses(document)

	for _, eq := range eqs {
		eq.Print()
	}
}
