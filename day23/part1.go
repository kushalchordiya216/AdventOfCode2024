package day23

import (
	"sort"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

type Part1Solver struct {
	graph map[string]mapset.Set[string]
}

func (s *Part1Solver) Read(path string) error {
	graph, err := ReadGraphFromFile(path)
	if err != nil {
		return err
	}
	s.graph = graph
	return nil
}

func (s *Part1Solver) Solve() int {
	triangles := mapset.NewSet[string]()
	
	// Find all nodes starting with 't'
	for tNode := range s.graph {
		if !strings.HasPrefix(tNode, "t") {
			continue
		}
		
		// For each neighbor of t-node
		tNeighbors := s.graph[tNode]
		for neighbor := range tNeighbors.Iter() {
			// Get neighbors of this neighbor
			nNeighbors := s.graph[neighbor]
			
			// Find intersection: nodes connected to both t-node and neighbor
			intersection := tNeighbors.Intersect(nNeighbors)
			
			// Each node in intersection forms a triangle with t-node and neighbor
			for thirdNode := range intersection.Iter() {
				// Create canonical triangle string (sorted order)
				triangle := []string{tNode, neighbor, thirdNode}
				sort.Strings(triangle)
				triangleKey := strings.Join(triangle, "-")
				
				triangles.Add(triangleKey)
			}
		}
	}
	
	return triangles.Cardinality()
}
