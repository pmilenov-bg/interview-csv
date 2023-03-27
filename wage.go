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
	// id	int
	name string 'scv: "name"'
	rate int	'scv: "rate"'
	hour int	'scv: "hore"'
	wage int	'scv: "-"'	//The "-" tag on the Wage field means that this field should be ignored when parsing the csv
}

func (s *State) calculateWage() {
	s.wage = s.rate * s.hour
	fmt.Printf("the wage of %s is %d\n", s.name, s.wage)
}


func main() {
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
