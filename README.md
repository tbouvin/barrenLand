# barrenLand
Install golang:
Go to web page and install golang for your platform: https://golang.org/dl/

To run:
From command line, run: go run barrenLand.go <Barren Land Coordinates>

Example: go run barrenLand.go {"48 192 351 207", "48 392 351 407",
  "120 52 135 547", "260 52 275 1000"}

Output: Sorted list of fertile regions in square meters

Running unit tests:
From command line, run: go test

Expected ouput from tests:
  Error validating coordinates INVCOORD
  Error validating coordinates INVCOORD
  Error validating coordinates OOBXS
  Error validating coordinates OOBYS
  Error validating coordinates OOBXE
  Error validating coordinates OOBYE
  Error validating coordinates INVCOORD
  116800 116800
  22816 192608
  240000
  Error validating coordinates OOBYE
  Error parsing arguments [48 192 351 207 48 392 351 407 120 52 135 547 260 52 275 1000] (OOBYE)
  PASS
  coverage: 99.0% of statements
  ok  	_/Users/tbouvin/barrenLand	0.997s
