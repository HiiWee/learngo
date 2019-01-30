// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	// parse the data
	rows := strings.Split(data, "\n")
	cols := strings.Split(rows[0], separator)

	// default case: slice for all the columns
	from, to := 0, len(cols)

	// find the from and to positions depending on the command-line arguments
	args := os.Args[1:]
	for i, v := range cols {
		l := len(args)

		if l >= 1 && v == args[0] {
			from = i
		}

		if l == 2 && v == args[1] {
			to = i + 1 // +1 because the stopping index is a position
		}
	}

	for i, row := range rows {
		cols := strings.Split(row, separator)

		// print the only the requested columns
		for _, h := range cols[from:to] {
			fmt.Printf("%-15s", h)
		}
		fmt.Println()

		// print extra new line for the header
		if i == 0 {
			fmt.Println()
		}
	}
}
