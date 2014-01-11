/*
 * provides adjacency list based implementation of undirected
 * graphs
**/
package graph

import (
	"container/list"
)

//
// adjacency list representation of a Graph, which contains 'V'
// vertices and 'E' edges. vertices are in the range {0, V-1} for ease
// of processing
//
type Graph struct {
	v   int32
	e   int32
	adj []list.List
}

//
// this function is called to create a new skeleton graph, with a
// specific number of vertices
//
func New(V int32) *Graph {
	return &Graph{
		v:   V,
		e:   0,
		adj: make([]list.List, V),
	}
}

func (G *Graph) V() int32 { return G.v }
func (G *Graph) E() int32 { return G.e }

//
// return the list of vertices adjacent to a given vertex 'v'.
//
func (G *Graph) Adj(v int32) []int32 {
	adj_list := G.adj[v]
	vertex_list := make([]int32, adj_list.Len())

	for i, node := 0, adj_list.Front(); node != nil; i, node = i+1, node.Next() {
		v := node.Value.(int32)
		vertex_list[i] = v
	}

	return vertex_list
}

//
// in a graph G, add an edge between vertices 'v' and 'w'. for
// undirected graphs, this operation adds v-w, and w-v edges as
// well
//
func (G *Graph) AddEdge(v, w int32) {
	G.adj[v].PushBack(w)
	G.adj[w].PushBack(v)

	G.e += 1

	return
}

//
// enumerate some fundamental properties of a graph
//
func (G *Graph) Degree(v int32) (degree int32) { return int32(len(G.Adj(v))) }
func (G *Graph) AverageDegree() float64        { return float64(2 * G.V() / G.E()) }

func (G *Graph) MaxDegree() (max_degree int32) {
	max_degree = 0

	for v := int32(0); v < G.V(); v++ {
		d := G.Degree(v)
		if d > max_degree {
			max_degree = d
		}
	}

	return
}
