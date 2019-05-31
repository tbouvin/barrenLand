/*
 * File: barrenLand.go
 * Author: Taylor Bouvin
 * Date: 5/31/19
 *
 * Barren Land Analysis
 * Objective - You have a farm of 400m by 600m where coordinates of the field are
 * from (0, 0) to (399, 599). A portion of the farm is barren, and all the barren
 * land is in the form of rectangles. Due to these rectangles of barren land, the
 * remaining area of fertile land is in no particular shape. An area of fertile
 * land is defined as the largest area of land that is not covered by any of the
 * rectangles of barren land. Read input from STDIN. Print output to STDOUT
 * Implement a Breadth-First Search to traverse the grid and find all sections
 * of fertile land.
 *
 * Think of the grid as a disconnected graph.
 * Each fertile vertex is connected to its adjacent fertile vertices.
 * Each barren vertex is connected to its adjacent barren vertices.
 * A barren vertex CANNOT be connected to a fertile vertice and vice-versa
 *
 * Because the graph is disconnected, we need to ensure that all vertices are
 * traversed. To do this, the main body should be enclosed in a loop that
 * starts at (0,0) and checks if that vertex has been visited. If not, then
 * perform BFS starting at that vertex. If it has been visited, then continue to
 * the next vertex in the list. Each loop of the BFS function will find a different
 * fertile regions area.
 *
 * When performing BFS for a particular vertex, the follwoing should be performed
 *	while vertices in queue
 *		pop vertex off queue
 *  	increment area for current BFS loop
 *  	for each adjacent neighbor:
 *			add neighbor to queue
 *			mark neighbor as visited
 * 	return area
 *
 * We only need to calculate each area of fertile land. This means that all barren
 * areas can be marked as visited during construction of the grid.
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// UPPERX - Upper bounds of X-coordinate of farm
const UPPERX = 400

// UPPERY - Upper bounds of Y-coordinate of farm
const UPPERY = 600

// Vertex - struct to represent x-y coordinates on grid
type Vertex struct {
	x, y int
}

// ExcludedArea - struct defining bottom left x-y corrdinates of barren land (s)
// and top right x-y coordinates of barren land (e)
type ExcludedArea struct {
	start, end Vertex
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Not enough arguments\n")
		return
	}

	_ = FindFertileLand(os.Args[1:])
}

// FindFertileLand - Find the area of fertile land in for a farm that has patches
// of barren land. Barren land coordinates are provided as a string array
// in the arguments. Each barren land region is defined by x-y coordinates for the
// lower left point and the upper right point.
//
// The area of each fertile region is returned in a sorted array from small to large.
func FindFertileLand(args []string) []int {
	ex, err := ParseArgs(args)
	if err != nil {
		fmt.Printf("Error parsing arguments %s (%s)\n", args, err.Error())
		return nil
	}

	land := BFS(ex)
	for v := range land {
		fmt.Printf("%d ", land[v])
	}

	fmt.Printf("\n")

	return land
}

// Delimiters -list of runes that the input string should be split on. If the provided rune
// matches any of the defined runes 'true' is returned
func Delimiters(r rune) bool {
	return (r == '{' || r == '}' || r == ',' || r == ' ' || r == '"' || r == '”' || r == '“')
}

// ParseArgs -Parses the input string and returns a ExcludedArea array that defines all
// rectangular barren regions.
// nil is returned if input string is invalid
func ParseArgs(parts []string) ([]ExcludedArea, error) {
	newString := strings.Join(parts, " ")
	parsedString := strings.FieldsFunc(newString, Delimiters)
	parsedStringLen := len(parsedString)
	if parsedStringLen%4 != 0 {
		return nil, errors.New("INVCOORDCNT")
	}

	var excluded []ExcludedArea
	for i := 0; i < parsedStringLen; i += 4 {
		unvalidatedCoords := [4]string{
			parsedString[i],
			parsedString[i+1],
			parsedString[i+2],
			parsedString[i+3]}

		validCoords, err := ValidateCoordinates(unvalidatedCoords)
		if err != nil {
			fmt.Printf("Error validating coordinates %s\n", err.Error())
			return nil, err
		}

		excluded = append(excluded, validCoords)
	}

	return excluded, nil
}

// ValidateCoordinates - Verifies that coordinates passed as a string are valid
// integers and within the bounds of the grid
func ValidateCoordinates(unvalidCoords [4]string) (ExcludedArea, error) {
	var coord ExcludedArea
	var xserr, yserr, xeerr, yeerr error
	coord.start.x, xserr = strconv.Atoi(unvalidCoords[0])
	coord.start.y, yserr = strconv.Atoi(unvalidCoords[1])
	coord.end.x, xeerr = strconv.Atoi(unvalidCoords[2])
	coord.end.y, yeerr = strconv.Atoi(unvalidCoords[3])

	if xserr != nil || yserr != nil || xeerr != nil || yeerr != nil {
		return coord, errors.New("INVCOORD")
	}

	if coord.start.x < 0 || coord.start.x >= UPPERX {
		return coord, errors.New("OOBXS")
	}

	if coord.start.y < 0 || coord.start.y >= UPPERY {
		return coord, errors.New("OOBYS")
	}

	if coord.end.x < 0 || coord.end.x >= UPPERX {
		return coord, errors.New("OOBXE")
	}

	if coord.end.y < 0 || coord.end.y >= UPPERY {
		return coord, errors.New("OOBYE")
	}

	return coord, nil
}

func isVertexBarren(p Vertex, ex []ExcludedArea) bool {
	//Check if vertex is barren
	for _, i := range ex {
		if p.x >= i.start.x && p.x <= i.end.x &&
			p.y >= i.start.y && p.y <= i.end.y {
			// It is barren
			return true
		}
	}

	return false
}

func isVertexInbounds(p Vertex) bool {
	if p.x >= 0 && p.x < UPPERX && p.y >= 0 && p.y < UPPERY {
		return true
	}

	return false
}

// GetAdjacentVertices - Retrieve all of the adjacent vertices for the given
// vertex. Adjacent vertices are returned in an array of vertices
func GetAdjacentVertices(p Vertex, ex []ExcludedArea) []Vertex {
	var adj []Vertex
	barren := isVertexBarren(p, ex)

	rightAdj := Vertex{x: p.x + 1, y: p.y}
	if isVertexInbounds(rightAdj) && (isVertexBarren(rightAdj, ex) == barren) {
		adj = append(adj, rightAdj)
	}

	leftAdj := Vertex{x: p.x - 1, y: p.y}
	if isVertexInbounds(leftAdj) && (isVertexBarren(rightAdj, ex) == barren) {
		adj = append(adj, leftAdj)
	}

	upAdj := Vertex{x: p.x, y: p.y + 1}
	if isVertexInbounds(upAdj) && (isVertexBarren(upAdj, ex) == barren) {
		adj = append(adj, upAdj)
	}

	downAdj := Vertex{x: p.x, y: p.y - 1}
	if isVertexInbounds(downAdj) && (isVertexBarren(downAdj, ex) == barren) {
		adj = append(adj, downAdj)
	}

	return adj
}

// ConstructMap - Builds an adjacency map (adj) with a key for each x-y vertex and
// it's value is all adjacent neighbors. Neighbor are adjacent if they share the
// same state (Barren/Fertile)
//
// A visited map (v) is also constructed with a key for each x-y vertex and it's
// value is if the vertex has been visited by BFS.
// NOTE: a vertex is marked as visited if it is in an excluded region
func ConstructMap(ex []ExcludedArea) (map[Vertex][]Vertex, map[Vertex]bool) {
	// Create necessary maps
	var adj map[Vertex][]Vertex
	var v map[Vertex]bool
	adj = make(map[Vertex][]Vertex)
	v = make(map[Vertex]bool)

	for i := 0; i < UPPERX; i++ {
		for j := 0; j < UPPERY; j++ {
			cur := Vertex{x: i, y: j}
			// Find all adjacent vertices for the current vertex and pass them to
			// the adjacency map
			adj[cur] = GetAdjacentVertices(cur, ex)

			// The barren areas don't need to be traversed, so mark as visited
			// from the get-go
			if isVertexBarren(cur, ex) {
				v[cur] = true
			} else {
				v[cur] = false
			}
		}
	}

	return adj, v
}

// BFSLoop - Every time this runs, it retrieves an area of fertile land
// The loop will find adjacent neighbors and increment the area until there
// are no more adjacent neighbors.
func BFSLoop(s Vertex, adj map[Vertex][]Vertex, v map[Vertex]bool, ex []ExcludedArea) int {
	area := 0
	// Set the starting vertex to visited
	v[s] = true
	var q []Vertex
	// Add the starting vertex to the queue
	q = append(q, s)
	// While the queue is not empty, perform a search
	for len(q) > 0 {
		// get vertex at front of queue
		cur := q[0]
		// pop off queue
		q = q[1:]
		// Another vertex has been added to this area, increment accordingly
		area++
		// Get adjacent vertices for current vertex
		neigh := adj[cur]
		for i := 0; i < len(neigh); i++ {
			// If neighbor has not been visited, then mark as visited and push to queue
			if !v[neigh[i]] {
				v[neigh[i]] = true
				q = append(q, neigh[i])
			}
		}
	}

	return area
}

// BFS - Entry point for the BFS algorithm, BFS will attempt to perform on each
// unvisited vertex. This ensures that even if the grid is disconnected that all
// regions will be traversed (besides excluded regions)
func BFS(ex []ExcludedArea) []int {
	var areaList []int
	// Construct maps needed for BFS
	m, v := ConstructMap(ex)
	for i := 0; i < UPPERX; i++ {
		for j := 0; j < UPPERY; j++ {
			cur := Vertex{x: i, y: j}
			// Only perform BFS on vertices that have not been visited. If a vertex
			// has been visited, then its area has already been accounted for.
			if !v[cur] {
				areaList = append(areaList, BFSLoop(cur, m, v, ex))
			}
		}
	}

	// sort the list from smallest to largest
	sort.Ints(areaList)
	return areaList
}
