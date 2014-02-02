package graph

//
// a set of typically used operations on graphs
//
type GraphOps interface {
	V() int32          // number of vertices
	Adj(int32) []int32 // adjacency list
}
