package main

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	sum := Calculator(7, 7)
	if sum != 7 {
		t.Error("7 was expected and we obtained", sum)
	}

	fmt.Println("Test of successful sum")
}

func TestConditionals(t *testing.T) {
	x := Conditionals(7)
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
}
