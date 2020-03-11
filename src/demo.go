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

	// Dataset
	/*
		var dataset anonmodel.Dataset
		var field1 = anonmodel.FieldAnonymizationInfo{Name: "City", Mode: "cat"} // Kategorikus
		var field2 = anonmodel.FieldAnonymizationInfo{Name: "Age", Mode: "int"}  // Intervallum
		var fields []anonmodel.FieldAnonymizationInfo
		dataset.Fields = append(fields, field1, field2)
	*/

	// Document
	var document = make(map[string]interface{})
	document["City"] = "Bp"
	document["Age"] = anonmodel.Interval{40, 50}

	var eqs = anonbll.GetMatchingClasses(document)

	for _, eq := range eqs {
		eq.Print()
	}
}

func CreateDemoEqulivalenceClasses() {

	var interval = make(map[string]anonmodel.Interval)
	interval["Age"] = anonmodel.Interval{10, 20}

	var categoric = make(map[string]string)
	categoric["City"] = "Bp"

	var c1 = anonmodel.EqulivalenceClass{1, categoric, interval, 0, true}

	anondb.CreateEqulivalenceClass(&c1)
}
