package main

import "fmt"
func main() {
	fmt.Println("Project started successfully!")
	d := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	o := []int{1, 1, 0, 1, 0, 1, 1}
	s := []string{}
	g := []string{}
	i := 0
	for i = 0; i < len(d); i++ {
		if o[i] == 1 {
			s = append(s, d[i])
		} else if o[i] == 0 {
			g = append(g, d[i])
		}
	}
	fmt.Printf("%+v\n", d)
	fmt.Printf("%+v\n", o)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", g)
}
