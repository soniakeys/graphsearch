package adj_test

import (
	"fmt"
	"sort"

	"github.com/soniakeys/graph"
	"github.com/soniakeys/graph/adj"
)

func ExampleEdge_Distance() {
	// Example shows that adj.Edge implements graph.DistanceEdge.
	var e graph.DistanceEdge
	e = adj.Edge(4)
	fmt.Println(e.Distance())
	// Output:
	// 4
}

func ExampleGraph_Link() {
	g := adj.Graph{}

	// As a minimimal example, use ints for nodes and don't use edges at all.
	g.Link(1, 2, nil)
	g.Link(2, 3, nil)
	g.Link(2, 1, nil)

	// Buffer and sort output because maps are unordered.
	var output []string
	for id, nd := range g {
		// For each node, print the node.
		line := fmt.Sprintf("node %v: neighbors:", id)
		nd.Visit(func(nb graph.Neighbor) {
			// Print a list of neighbors on the same line.
			line += fmt.Sprintf(" %v", nb.Nd)
		})
		output = append(output, line)
	}
	sort.Strings(output)
	for _, line := range output {
		fmt.Println(line)
	}
	// Output:
	// node 1: neighbors: 2
	// node 2: neighbors: 3 1
	// node 3: neighbors:
}