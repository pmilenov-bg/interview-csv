package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
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
func main() {
	// #1 open a file
	f, err := os.Open("empl.csv")
	if err != nil {
		// fmt.Println(err)
		log.Fatalln(err)
	}
	defer f.Close()
	// #2 perse a csv file
	csvReader := csv.NewReader(f)
	for rowCount := 0; ; rowCount++ {
		r, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		// columns := make(map[string]int)
		if rowCount == 0 {
			fmt.Println(";lkasfdl")
			// for
		}
		fmt.Println(r)

	}
	// #3 check for 0
	// #4 check for repeated Names
	// #5 wage calculation

}
