package main

import "fmt"

func main() {
	// Basic math stuff
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// Basic for loop
	a := 1
	if a > 6 {
		fmt.Println("more than six")
	} else if a > 2 {
		fmt.Println("More than two but less than six")
	} else {
		fmt.Println("less than two")
	}

	// Basic array //

	// long hand array init
	var b [5]int
	b[0] = 1
	b[1] = 6
	fmt.Println(b)

	// short hand array int
	c := [5]int{5, 4, 3, 2, 1}
	fmt.Println(c)

	// slice of ints
	d := []int{6, 7, 8, 9, 10}
	d = append(d, 12)

	fmt.Println(d)

}
