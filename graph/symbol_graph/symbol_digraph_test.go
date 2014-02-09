package symbol_graph

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph/algorithms"
)

// this function tests whether the topological-sort for a digraph
// matches the expected output
func ExampleTopoOrder() {
	fname := "../data/jobs-symgraph.data"

	// load it
	sym_dg, err := DigraphFromFile(fname, "/")
	if err != nil {
		panic(err)
	}

	// compute the ordering
	topo_order, err := algorithms.ComputeTopologicalOrder(sym_dg.G())
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	for _, v := range topo_order.Order() {
		vname, _ := sym_dg.Name(v)
		fmt.Printf("%s\n", vname)
	}

	// Output:
	// Calculus
	// Linear Algebra
	// Introduction to CS
	// Advanced Programming
	// Algorithms
	// Theoretical CS
	// Artificial Intelligence
	// Robotics
	// Machine Learning
	// Neural Networks
	// Databases
	// Scientific Computing
	// Computational Biology
}
