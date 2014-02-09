package algorithms

import (
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/traversal"
)

type Topological struct {
	order []int32
}

//
// this function computes the topological order for a given
// digraph. returns an error if no such ordering is possible
//
func ComputeTopologicalOrder(G graph.GraphOps) (TO *Topological, err error) {
	if dg := DigraphCyle(G); !dg.IsAcyclic() {
		err = fmt.Errorf("error: given no ordering possible. digraph has cycle: %v\n", dg.Cycle())
		return
	}

	//
	// reverse-post-order traversal is the topological sort
	// order.
	//
	dg_to := traversal.DoDFSTraversals(G)
	topo_order := dg_to.ReversePost()

	TO = &Topological{
		order: make([]int32, len(topo_order)),
	}

	copy(TO.order, topo_order)

	return
}

func (TO *Topological) Order() []int32 { return TO.order }
func (TO *Topological) IsDAG() bool    { return len(TO.order) != 0 }
