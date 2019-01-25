package main

import (
	"fmt"
	"strings"
)

/// restaurant project
func main() {
	fmt.Println("Project started successfully!")
	// week days
	d := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	//*** 0 for closing days and 1 for opening days ***//
	// o := []int{0, 1, 0, 0, 1, 0, 1}
	o := []int{0, 0, 0, 0, 0, 0, 0}
	// o := []int{1, 1, 1, 1, 1, 1, 1}
	// o := []int{0, 1, 0, 0, 0, 0, 1}
	// o := []int{1, 0, 0, 0, 1, 0, 1}
	// o := []int{0, 0, 0, 0, 1, 0, 1}
	// o := []int{1, 1, 1, 0, 1, 0, 1}
	// o := []int{0, 0, 0, 0, 0, 1, 0}
	var s []string
	var g []string
	var start int
	var i int
	var j int
	//for/while loop to loop through until aaray of o ends
	for i < len(d) {
		if o[i] == 1 {
			start = i // start of opening days
			j = i + 1 // next day after start
			//checking condition for opening days until week ends
			if j < len(d) && o[j] == 1 {
				for j < len(d) && o[j] != 0 {
					i++
					j++
				}
				p := d[start] + "-"
				p = p + d[i]
				s = append(s, p) // apending start day to the array
			} else {
				s = append(s, d[start]) // apending start day to the array if only one open day is there
			}
		} else { //checking condition for closing days until week ends
			start = i // start of closing days
			j = i + 1 // next day after start
			if j < len(d) && o[j] == 0 {
				for j < len(d) && o[j] != 1 {
					i++
					j++
				}
				f := d[start] + "-" + d[i]
				g = append(g, f) // apending start day to the array

			} else {
				g = append(g, d[start]) // apending start day to the array if only one open day is there
			}
		}
		i++
	}
	fmt.Printf("%+v\n", d)                       // array of  week days
	fmt.Printf("%+v\n", o)                       // array of  opening and closing days
	fmt.Printf("%+v\n", strings.Join(s[:], ",")) // array of opening days
	fmt.Printf("%+v\n", strings.Join(g[:], ",")) // array of closing days
}
