package main

import (
	"encoding/csv"
	"fmt"
	"io"

	// "io"
	"log"
	"os"
	"strconv"
)

type State struct {
	name string `csv:"name"`
	rate int    `csv:"rate"`
	hour int    `csv:"hour"`
	wage int    `csv:"-"`
}

func (s *State) calculateWage() {
	s.wage = s.rate * s.hour
	fmt.Printf("the wage of %s is %d\n", s.name, s.wage)
}

func readCsvFile(filePath string) []*State {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file " + filePath)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + filePath)
	}

	var states []*State

	for _, record := range records {
		state := &State{
			name: record[0],
			rate: atoi(record[1]),
			hour: atoi(record[2]),
		}
		states = append(states, state)
	}

	return states
}

// func testInputs() {
// 	// test for repetitive names
// 	// test for zero values -> output to the screen
// 	// test for negative values
// 	// test the field for type inconsistency
// }

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Unable to convert %s to int", s)
	}
	return n
}

func main() {
	in := readCsvFile

	// #1 open a file
	f, err := os.Open("input1.csv")
	if err != nil {
		// fmt.Println(err)
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 perse a csv file
	csvReader := csv.NewReader(f)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("----")
		fmt.Println("Raw data from  the csv file", row)

		rate, err := strconv.Atoi(row[1])
		if err != nil {
			// ... handle error
			panic(err)
		}

		hour, err := strconv.Atoi(row[2])
		if err != nil {
			// ... handle error
			panic(err)
		}

		if hour == 0 {
			fmt.Println("skip the line")
			continue // discontinue the cycle started line 46
		}
	}
}
