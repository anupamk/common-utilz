package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/traversal"
)

//
// this function computes the topological order for a given
// digraph. returns an error if no such ordering is possible
//
func ComputeTopologicalOrder(G graph.GraphOps) (ordering []int32, err error) {
	if yes, cycle := IsDigraphAcyclic(G); yes {
		err = fmt.Errorf("error: digraph has a cycle: %v. no ordering possible\n", cycle)
		return
	}

	//
	// reverse-post-order traversal is the topological sort
	// order.
	//
	dg_to := traversal.DoDFSTraversals(G)
	topo_order := dg_to.ReversePost()
	ordering = make([]int32, len(topo_order))

	// copy the result
	copy(ordering, topo_order)

	return
}
