package main

import "fmt"

var d = []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
var o = []int{1, 1, 1, 0, 1, 0, 0}
var one string
var zero string
var p = 0
var i, k, j int
var b, c string

func main() {

	for p < len(d) {
		if o[i] == 1 {
			k, one := sort(i, 1)
			i = k
			fmt.Println("one->", one)
		} else {
			k, zero := sort(i, 0)
			i = k
			fmt.Println("zero->", zero)
		}
		i += 1
	}
	fmt.Println("one->", one)
	fmt.Println("zero->", zero)

}
func sort(i int, k int) (index int, a string) {
	var start int
	var j int
	start = i
	j = i + 1

	if o[j] == k {
		for o[j] == k && o[j] <= len(o) {
			i++
			j++
		}
		b = d[start] + d[i]
		return i, b
	} else {
		c = d[start]
		return i, c
	}

}
