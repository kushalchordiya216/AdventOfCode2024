package day23

import (
	"fmt"
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Part2Solver struct {
	graph     map[string]mapset.Set[string]
	maxClique mapset.Set[string]
}

func (s *Part2Solver) Read(path string) error {
	graph, err := ReadGraphFromFile(path)
	if err != nil {
		return err
	}
	s.graph = graph
	s.maxClique = mapset.NewSet[string]()
	return nil
}

func (s *Part2Solver) Solve() int {
	// Get all nodes
	allNodes := mapset.NewSet[string]()
	for node := range s.graph {
		allNodes.Add(node)
	}

	// Start Bron-Kerbosch algorithm
	R := mapset.NewSet[string]() // Current clique
	P := allNodes.Clone()        // Candidate nodes
	X := mapset.NewSet[string]() // Exclusion set

	s.bronKerbosch(R, P, X)

	// Return the nodes in sorted order as requested
	clique := s.maxClique.ToSlice()
	sort.Strings(clique)

	// For AoC, we typically need to return an integer
	// But since you want the actual nodes, we'll print them and return size
	fmt.Println(strings.Join(clique, ","))
	return len(clique)
}

func (s *Part2Solver) bronKerbosch(R, P, X mapset.Set[string]) {
	// Base case: if P and X are both empty
	if P.Cardinality() == 0 && X.Cardinality() == 0 {
		// Check if this clique has at least one 't' node
		hasTNode := false
		for node := range R.Iter() {
			if strings.HasPrefix(node, "t") {
				hasTNode = true
				break
			}
		}

		// Update max clique if this one is larger and has 't' node
		if hasTNode && R.Cardinality() > s.maxClique.Cardinality() {
			s.maxClique = R.Clone()
		}
		return
	}

	// Choose pivot to minimize recursive calls
	// Pivot should be the node with most neighbors in P
	pivot := s.choosePivot(P, X)

	// Create working copy of P for iteration
	PminusPivotNeighbors := P.Clone()
	if pivot != "" {
		pivotNeighbors := s.graph[pivot]
		PminusPivotNeighbors = P.Difference(pivotNeighbors)
	}

	// For each node in P \ neighbors(pivot)
	for node := range PminusPivotNeighbors.Iter() {
		nodeNeighbors := s.graph[node]

		// Pruning: if current clique + remaining candidates can't beat max, skip
		if R.Cardinality()+P.Cardinality() <= s.maxClique.Cardinality() {
			break
		}

		// Recursive call
		newR := R.Clone()
		newR.Add(node)
		newP := P.Intersect(nodeNeighbors)
		newX := X.Intersect(nodeNeighbors)

		s.bronKerbosch(newR, newP, newX)

		// Move node from P to X
		P.Remove(node)
		X.Add(node)
	}
}

func (s *Part2Solver) choosePivot(P, X mapset.Set[string]) string {
	PunionX := P.Union(X)
	bestPivot := ""
	maxConnections := -1

	// Choose pivot with maximum connections in P
	for candidate := range PunionX.Iter() {
		candidateNeighbors := s.graph[candidate]
		connections := P.Intersect(candidateNeighbors).Cardinality()

		if connections > maxConnections {
			maxConnections = connections
			bestPivot = candidate
		}
	}

	return bestPivot
}

// GetMaxCliqueNodes returns the nodes in the maximum clique for printing
func (s *Part2Solver) GetMaxCliqueNodes() []string {
	clique := s.maxClique.ToSlice()
	sort.Strings(clique)
	return clique
}
