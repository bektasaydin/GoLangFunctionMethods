package main

import (
	"fmt"
	"testing"
)

func TestIntegrationCalculator(t *testing.T) {
	calc := Calculator(8, 8)
	average := Average(8, 8)

	iototalaverage := Average(calc, average)

	if iototalaverage != 8 {
		t.Error("expected 17 but obtained", iototalaverage)
	}

	fmt.Println("Total average integration test was successful")
}
