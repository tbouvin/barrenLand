/*
 * File: barrenLand_test.go
 * Author: Taylor Bouvin
 * Date: 5/31/19
 */

package main

import (
	"testing"
)

func TestParseArgs(t *testing.T) {

	testString := []string{"48 192 351 207", "48 392 351 407",
		"120 52 135 547", "260 52 275 547"}
	parsedPairs, err := ParseArgs(testString)
	if err != nil {
		t.Errorf("Error parsing arguments")
	}

	expected := 4
	stringLen := len(parsedPairs)
	if stringLen != expected {
		t.Errorf("Number of parsed integers incorrect, got %d, want %d", stringLen, expected)
	}

	// Test that parser throws error when 4 coordinates are not given for the
	// barren regions
	invalidCoordNum := append(testString, "42")
	_, err = ParseArgs(invalidCoordNum)
	if err == nil {
		t.Errorf("Expected return value of nil")
	}

	// Test that parser throws error if a non-integer is provided in a x coordinates
	invalidXCoord := append(testString, "BADX 42 399 499")
	_, err = ParseArgs(invalidXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of non-integer X coordinate")
	}

	// Test that parser throws error if a non-integer is provided in a x coordinates
	invalidYCoord := append(testString, "42 BADY 399 499")
	_, err = ParseArgs(invalidYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of non-integer Y coordinate")
	}

	// Test that parser throws error if an out of bounds number is a x coordinate
	OOBXCoord := append(testString, "500 42 399 499")
	_, err = ParseArgs(OOBXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of OOB XS coordinate")
	}

	// Test that parser throws error if an out of bounds number is a y coordinate
	OOBYCoord := append(testString, "42 700 399 499")
	_, err = ParseArgs(OOBYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of OOB YS coordinate")
	}

	// Test that parser throws error if a negative number is a x coordinate
	OOBXCoord = append(testString, "0 42 -1 499")
	_, err = ParseArgs(OOBXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of negative XE coordinate")
	}

	// Test that parser throws error if a negative number is a x coordinate
	OOBYCoord = append(testString, "42 0 399 -1")
	_, err = ParseArgs(OOBYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of negative YE coordinate")
	}

	// Test that parser throws error if a negative number is a x coordinate
	DoubleCoord := append(testString, "42 0 399 100.1")
	_, err = ParseArgs(DoubleCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of non-integer YE coordinate")
	}
}

// Test first case provided by case study instructions
func TestBFSCase1(t *testing.T) {
	testString := []string{"0 292 399 307"}
	expectedArea := []int{116800, 116800}
	area := FindFertileLand(testString)

	for i, v := range area {
		if v != expectedArea[i] {
			t.Errorf("Area calculated for test1 was not correct")
		}
	}
}

// Test second case provided by case study instructions
func TestBFSCase2(t *testing.T) {
	testString := []string{"48 192 351 207", "48 392 351 407",
		"120 52 135 547", "260 52 275 547"}
	expectedArea := []int{22816, 192608}
	area := FindFertileLand(testString)

	for i, v := range area {
		if v != expectedArea[i] {
			t.Errorf("Area calculated for test1 was not correct")
		}
	}
}

// Test invalid coordinates
func TestBFSInvalid(t *testing.T) {
	testString := []string{"48 192 351 207", "48 392 351 407",
		"120 52 135 547", "260 52 275 1000"}
	area := FindFertileLand(testString)
	if area != nil {
		t.Errorf("Invalid coordinates should have caused FindFertileLands to return nil")
	}
}
