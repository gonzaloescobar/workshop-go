package main

import (
	"testing"
)

func TestAddDivisor0(t *testing.T) {

	expected := 0
	result := sumNumberDivisibleByThreeOrFive(0)

	if result != expected {
		t.Fatalf("I was expecting %d but I got %d", result, expected)
	}
}
func TestAddDivisor3(t *testing.T) {

	expected := 3
	result := sumNumberDivisibleByThreeOrFive(3)

	if result != expected {
		t.Fatalf("I was expecting %d but I got %d", result, expected)
	}
}

func TestAddDivisor5(t *testing.T) {

	expected := 8
	result := sumNumberDivisibleByThreeOrFive(5)

	if result != expected {
		t.Fatalf("I was expecting %d but I got %d", result, expected)
	}
}

func TestAddDivisor6(t *testing.T) {

	expected := 14
	result := sumNumberDivisibleByThreeOrFive(6)

	if result != expected {
		t.Fatalf("I was expecting %d but I got %d", result, expected)
	}
}

func TestAddDivisor9(t *testing.T) {

	expected := 23
	result := sumNumberDivisibleByThreeOrFive(9)

	if result != expected {
		t.Fatalf("I was expecting %d but I got %d", result, expected)
	}
}
