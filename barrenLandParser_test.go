package main

import (
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	testString := []string{"0 292 399 307", "0 292 399 307"}
	parsedPairs := ParseArgs(testString)
	var expected = 2
	var stringLen = len(parsedPairs)
	if stringLen != expected {
		t.Errorf("Number of parsed integers incorrect, got %d, want %d", stringLen, expected)
	}

	invalidCoordNum := append(testString, "42")
	if ParseArgs(invalidCoordNum) != nil {
		t.Errorf("Expected return value of nil")
	}

	invalidXCoord := append(testString, "BADX 42 399 499")
	if ParseArgs(invalidXCoord) != nil {
		t.Errorf("Expected return value of nil because of non-integer X coordinate")
	}

	invalidYCoord := append(testString, "42 BADY 399 499")
	if ParseArgs(invalidYCoord) != nil {
		t.Errorf("Expected return value of nil because of non-integer Y coordinate")
	}

	OOBXCoord := append(testString, "500 42 399 499")
	if ParseArgs(OOBXCoord) != nil {
		t.Errorf("Expected return value of nil because of OOB X coordinate")
	}

	OOBYCoord := append(testString, "42 700 399 499")
	if ParseArgs(OOBYCoord) != nil {
		t.Errorf("Expected return value of nil because of OOB Y coordinate")
	}

	OOBXCoord = append(testString, "-1 42 399 499")
	if ParseArgs(OOBXCoord) != nil {
		t.Errorf("Expected return value of nil because of OOB X coordinate")
	}

	OOBYCoord = append(testString, "42 -1 399 499")
	if ParseArgs(OOBYCoord) != nil {
		t.Errorf("Expected return value of nil because of OOB Y coordinate")
	}
}

func TestDelimiters(t *testing.T) {
	testDelimiters := []rune{'{', '}', ',', ' ', '"'}
	for i := 0; i < len(testDelimiters); i++ {
		if Delimiters(testDelimiters[i]) == false {
			t.Errorf("Delimiter for %c did not return true", testDelimiters[i])
		}
	}
}

func TestBFS(t *testing.T) {
	testString := []string{"0 292 399 307"}
	expectedArea := []int{116800, 116800}
	bl := ParseArgs(testString)
	area := BFS(bl)
	for i, v := range area {
		if v != expectedArea[i] {
			t.Errorf("Area calculated for test1 was not correct")
		}
	}
	fmt.Printf("Area: %d\n", area)

	testString = []string{"48 192 351 207", "48 392 351 407",
		"120 52 135 547", "260 52 275 547"}
	expectedArea = []int{22816, 192608}
	bl = ParseArgs(testString)
	area = BFS(bl)
	for i, v := range area {
		if v != expectedArea[i] {
			t.Errorf("Area calculated for test1 was not correct")
		}
	}
	fmt.Printf("Area: %d\n", area)
}
