package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total Records: ", len(rows)-1)
	fmt.Println("Air Temp: ", mean(rows, 1), median(rows, 1))
	fmt.Println("Barometric: ", mean(rows, 2), median(rows, 2))
	fmt.Println("Wind Speed: ", mean(rows, 7), median(rows, 7))
}

func median(rows [][]string, idx int) float64 {
	var sorted []float64

	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	sort.Float64s(sorted)

	if len(sorted)%2 == 0 {
		return 0.0
	}
	return 1.0
}

func mean(rows [][]string, idx int) float64 {
	var total float64
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}
