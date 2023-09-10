package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go with example")

	/*
		for {
			fmt.Println("loop")
			break
		}

		switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
		}

		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("It's before noon")
		default:
			fmt.Println("It's after noon")
		}

		The whatAmI function takes an interface parameter i.
		  Inside the function, there is a type switch (switch t := i.(type))
		  that checks the actual type of the interface value i.
		whatAmI := func(i interface{}) {
			switch t := i.(type) {
			case bool:
				fmt.Println("I'm a bool")
			case int:
				fmt.Println("I'm an int")
			default:
				fmt.Printf("Don't know type %T\n", t)
			}
		}
		whatAmI(true)
		whatAmI(1)
		whatAmI("hey")
	*/

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

}
