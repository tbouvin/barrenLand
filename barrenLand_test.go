package main

import (
	"fmt"
	"testing"
)

func TestParseArgs(t *testing.T) {
	testString := []string{"0 292 399 307", "0 292 399 307"}
	parsedPairs, err := ParseArgs(testString)
	var expected = 2
	var stringLen = len(parsedPairs)
	if stringLen != expected {
		t.Errorf("Number of parsed integers incorrect, got %d, want %d", stringLen, expected)
	}

	invalidCoordNum := append(testString, "42")
	_, err = ParseArgs(invalidCoordNum)
	if err == nil {
		t.Errorf("Expected return value of nil")
	}

	invalidXCoord := append(testString, "BADX 42 399 499")
	_, err = ParseArgs(invalidXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of non-integer X coordinate")
	}

	invalidYCoord := append(testString, "42 BADY 399 499")
	_, err = ParseArgs(invalidYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of non-integer Y coordinate")
	}

	OOBXCoord := append(testString, "500 42 399 499")
	_, err = ParseArgs(OOBXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of OOB XS coordinate")
	}

	OOBYCoord := append(testString, "42 700 399 499")
	_, err = ParseArgs(OOBYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of OOB YS coordinate")
	}

	OOBXCoord = append(testString, "0 42 -1 499")
	_, err = ParseArgs(OOBXCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of negative XE coordinate")
	}

	OOBYCoord = append(testString, "42 0 399 -1")
	_, err = ParseArgs(OOBYCoord)
	if err == nil {
		t.Errorf("Expected return value of nil because of negative YE coordinate")
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
	bl, err := ParseArgs(testString)
	if err != nil {
		t.Errorf("Unable to parse test string (%s)", err.Error())
	}

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
	bl, err = ParseArgs(testString)
	if err != nil {
		t.Errorf("Unable to parse test string (%s)", err.Error())
	}

	area = BFS(bl)
	for i, v := range area {
		if v != expectedArea[i] {
			t.Errorf("Area calculated for test1 was not correct")
		}
	}
	fmt.Printf("Area: %d\n", area)
}
