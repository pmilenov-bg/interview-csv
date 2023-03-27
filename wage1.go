package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type State struct {
	name string `csv:"name`
	rate int    `csv:"rate"`
	hour int    `csv:"hore"`
	wage int    `csv:"-"` // The "-" tag on the Wage field means that this field should be ignored when parsing the cs`
}

func (s *State) calculateWage() {
	s.wage = s.rate * s.hour
	fmt.Printf("the wage of %s is %d\n", s.name, s.wage)
}

func main() {
	f, err := os.Open("input1.csv")
	if err != nil {
		// fmt.Println(err)
		log.Fatalln(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f) // #2 perse a csv file
	csvReader.FieldsPerRecords = -1

	var states []*State
	// Use a slice of pointers to State:
	// Instead of using a slice of State, you can use a slice
	// of pointers to State. This would make your code more
	// efficient when you're appending elements to the slice,
	// as you wouldn't need to create a copy of the entire slice every time.

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

		if wage <= 1000 {
			fmt.Println("wage under 1000")
		}
		var single_state State
		single_state = State{name: row[0], rate: rate, hour: hour}
		// var single_state State = State{name: row[0], rate: rate, hour: hour}
		// single_state := State{name: row[0], rate: rate, hour: hour}
		// single_state.wage = Wage(rate, hour)
		single_state.calculateWage()
		// single_state = append(single_state, single_state.Wage)
		fmt.Println(single_state)

		// all_states = append(all_states, single_state)
		// fmt.Println(all_states)

	}

}
