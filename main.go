// Calculate the mean of two numbers
package main

import (
	"fmt"
)

func main() {

	result := Calculator(1, 2)
	println(result)

	println("------------------------------------")

	cond := Conditionals(6)
	println(cond)

	println("------------------------------------")

	exswitch := ExSwitch(2)
	println(exswitch)

	println("------------------------------------")

	exforloop := ExForloop(2)
	println(exforloop)

	println("------------------------------------")

	exstring := ExString("The colour of magic")
	println(exstring)

	println("------------------------------------")

	exstringformat := ExStringFormat("String Formating Example")
	println(exstringformat)

	println("------------------------------------")

}

//Calculate Function
func Calculator(x int, y int) int {

	fmt.Printf("x=%v, type of %T\n", x, x)
	fmt.Printf("y=%v, type of %T\n", y, y)

	var mean int = (x + y) / 2
	fmt.Printf("result: %v, type of %T\n", mean, mean)
	return mean
}

//For Conditionals
func Conditionals(x int) int {

	if x > 5 {
		fmt.Println("x is big")
	}

	if x > 20 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is very big")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}

	return x
}

func ExSwitch(x int) int {

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("many")
	}

	return int(x)
}

func ExForloop(x int) int {

	for i := 0; i < 3; i++ {
		fmt.Println(i + 1)
	}

	return x

}

func ExString(x string) string {

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("x[0] = %v (type %T)\n", x[0], x[0])
	fmt.Println(x[4:11])
	fmt.Println(x[4:])
	fmt.Println(x[:4])
	fmt.Println("t" + x[1:])

	return x
}

func ExStringFormat(x string) string {
	n := 42
	s := fmt.Sprintf("%d", n)
	fmt.Printf("x = %s (type %T)\n", s, s)

	return x
}
