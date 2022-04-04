package main

import (
	"fmt"

	"rsc.io/quote"
	// "example.com/greetings"
)

func main() {

	fmt.Println(quote.Glass())
	// Basic math stuff
	x := 5
	y := 7
	sum := x + y
	fmt.Println(sum)

	// Basic if else
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

	// slice of ints so you can add onto array
	d := []int{6, 7, 8, 9, 10}
	d = append(d, 12)
	fmt.Println(d)

	// maps
	verticles := make(map[string]int)

	verticles["triangle"] = 2
	verticles["square"] = 3
	verticles["circle"] = 11

	delete(verticles, "circle")
	fmt.Println(verticles["square"])

	// go for loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// go while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// range loop
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	// message := greetings.Hello("Nick")
	// fmt.Println(message)
}
