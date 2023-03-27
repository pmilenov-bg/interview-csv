package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type State struct {
	name string `csv:"name`
	rate int    `csv:"rate"`
	hour int    `csv:"hore"`
	wage int    `csv:"-"` // The "-" tag on the Wage field means that this field should be ignored when parsing the cs`
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return records
}

// func testInputs() {
// 	// test for repetative names
// 	// test for zero values -> output to the screen
// 	// test for negative values
// 	// test the field for type inconsistancy
// }

func (s *State) calculateWage() {
	s.wage = s.rate * s.hour
	fmt.Printf("the wage of %s is %d\n", s.name, s.wage)
}


func main() {
	records := readCsvFile("./input.csv")
	for (record int records) {
		fmt.Println(record)
	}
	output := 

	// fmt.Println(records)
	// rawData = struct...{...}
	// results = struct ...{....}
	// read csv file into strct
	// print...
	// interate over the struct
	// for (row in rawData) {
	//    results.name = row.name
	//    results.wage = = raw[wage] * raw[hours]
	// }
	//
	// output the new struct
	// for (result in results)
	// 	  println result.name, result.wage
	// // You can edit this code!
	// Click here and start typing.

}
