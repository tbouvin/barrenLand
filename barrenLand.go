package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const FARMX = 400
const FARMY = 600

/*
 * Implement a Breadth-First Search to traverse the grid and find all sections
 * of fertile/barren land.
 *
 * Think of the grid as a connected graph.
 * Each fertile vertex is connected to its adjacent fertile vertices.
 * Each barren vertex is connected to its adjacent barren vertices.
 * A barren vertex CANNOT be connected to a fertile vertice and vice-versa
 *
 * Because the graph is disconnected, we need to ensure that all vertices are
 * traversed. To do this, the main body should be enclosed in a loop that
 * starts at (0,0) and checks if that vertex has been visited. If not, then
 * perform BFS starting at that vertex.
 */

func main() {
	args := os.Args[1:]
	bl := ParseArgs(args)

	land := BFS(bl)
	for v := range land {
		fmt.Printf("%d ", land[v])
	}

	fmt.Printf("\n")
}

func Delimiters(r rune) bool {
	return r == '{' || r == '}' || r == ',' || r == ' ' || r == '"'
}

type Point struct {
	x, y int
}

type BarrenLand struct {
	s, e Point
}

func ParseArgs(parts []string) []BarrenLand {
	newString := strings.Join(parts, " ")
	parsedString := strings.FieldsFunc(newString, Delimiters)
	parsedStringLen := len(parsedString)
	if parsedStringLen%4 != 0 {
		return nil
	}

	var barrenCoords []BarrenLand
	for i := 0; i < parsedStringLen; i += 4 {
		unvalidatedCoords := [4]string{
			parsedString[i],
			parsedString[i+1],
			parsedString[i+2],
			parsedString[i+3]}

		validCoords, err := ValidateCoordinates(unvalidatedCoords)
		if err != nil {
			fmt.Printf("Error validating coordinates %s\n", err.Error())
			return nil
		}

		barrenCoords = append(barrenCoords, validCoords)
	}

	return barrenCoords
}

func ValidateCoordinates(unvalidCoords [4]string) (BarrenLand, error) {
	var ret BarrenLand
	var xserr, yserr, xeerr, yeerr error
	ret.s.x, xserr = strconv.Atoi(unvalidCoords[0])
	ret.s.y, yserr = strconv.Atoi(unvalidCoords[1])
	ret.e.x, xeerr = strconv.Atoi(unvalidCoords[2])
	ret.e.y, yeerr = strconv.Atoi(unvalidCoords[3])

	if xserr != nil || yserr != nil || xeerr != nil || yeerr != nil {
		return ret, errors.New("INVCOORD")
	}

	if ret.s.x < 0 || ret.s.x >= FARMX {
		return ret, errors.New("INVCOORD")
	}

	if ret.s.y < 0 || ret.s.y >= FARMY {
		return ret, errors.New("INVCOORD")
	}

	return ret, nil
}

func isPointBarren(p Point, bl []BarrenLand) bool {
	//Check if point is barren
	for i := 0; i < len(bl); i++ {
		if p.x >= bl[i].s.x && p.x <= bl[i].e.x &&
			p.y >= bl[i].s.y && p.y <= bl[i].e.y {
			// It is barren
			return true
		}
	}

	return false
}

func isPointInbounds(p Point) bool {
	if p.x >= 0 && p.x < FARMX && p.y >= 0 && p.y < FARMY {
		return true
	}

	return false
}

func GetAdjacentPoints(p Point, bl []BarrenLand) []Point {
	var adj []Point
	barren := isPointBarren(p, bl)

	rightAdj := Point{x: p.x + 1, y: p.y}
	if isPointInbounds(rightAdj) && (isPointBarren(rightAdj, bl) == barren) {
		adj = append(adj, rightAdj)
	}

	leftAdj := Point{x: p.x - 1, y: p.y}
	if isPointInbounds(leftAdj) && (isPointBarren(rightAdj, bl) == barren) {
		adj = append(adj, leftAdj)
	}

	upAdj := Point{x: p.x, y: p.y + 1}
	if isPointInbounds(upAdj) && (isPointBarren(upAdj, bl) == barren) {
		adj = append(adj, upAdj)
	}

	downAdj := Point{x: p.x, y: p.y - 1}
	if isPointInbounds(downAdj) && (isPointBarren(downAdj, bl) == barren) {
		adj = append(adj, downAdj)
	}

	return adj
}

func ConstructMap(bl []BarrenLand) (map[Point][]Point, map[Point]bool) {
	var m map[Point][]Point
	var v map[Point]bool
	m = make(map[Point][]Point)
	v = make(map[Point]bool)

	for i := 0; i < FARMX; i++ {
		for j := 0; j < FARMY; j++ {
			cur := Point{x: i, y: j}
			m[cur] = GetAdjacentPoints(cur, bl)
			// fmt.Printf("Adj Point: (%d, %d): ", cur.x, cur.y)
			// fmt.Println(m[cur])

			//The barren areas don't need to be traversed, so mark as visited
			//from the get-go
			if isPointBarren(cur, bl) {
				v[cur] = true
			} else {
				v[cur] = false
			}
		}
	}

	return m, v
}

//Every time this runs, it fills out another spot of land
func BFSLoop(s Point, m map[Point][]Point, v map[Point]bool, bl []BarrenLand) int {
	area := 0
	v[s] = true
	var q []Point
	q = append(q, s)
	//While the queue is not empty, perform a search
	for len(q) > 0 {
		cur := q[0]
		q = q[1:] //pop
		// fmt.Printf("Point: (%d, %d)\n", cur.x, cur.y)
		area++
		//Get adjacent points for current point
		adj := m[cur]
		for i := 0; i < len(adj); i++ {
			if !v[adj[i]] {
				v[adj[i]] = true
				q = append(q, adj[i])
			}
		}
	}

	return area
}

func BFS(bl []BarrenLand) []int {
	var fertileLands []int
	m, v := ConstructMap(bl)
	// count := 0
	for i := 0; i < FARMX; i++ {
		for j := 0; j < FARMY; j++ {
			cur := Point{x: i, y: j}
			if !v[cur] {
				// fmt.Printf("Count: %d\n", count)
				// count++
				fertileLands = append(fertileLands, BFSLoop(cur, m, v, bl))
			}
		}
	}

	sort.Ints(fertileLands)
	return fertileLands
}
