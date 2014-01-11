package graph

import (
	"testing"
)

//
// definition of graph 'structure'
//
type graph_definition struct {
	v     int32
	e     int32
	edges [][2]int32
}

//
// properties of graphs that we care about
//
type graph_properties struct {
	average_degree float64
	maximum_degree int32
	vertex_degree  []int32
}

//
// graph definition and their properties in one place
//
type test_graph struct {
	graph_defn  graph_definition
	graph_props graph_properties
}

// a few test graphs
var graphs = [...]test_graph{

	// graph-0
	{
		graph_defn: graph_definition{
			v: 13,
			e: 13,
			edges: [][2]int32{
				{0, 5},
				{4, 3},
				{0, 1},
				{9, 12},
				{6, 4},
				{5, 4},
				{0, 2},
				{11, 12},
				{9, 10},
				{0, 6},
				{7, 8},
				{9, 11},
				{5, 3},
			},
		},

		graph_props: graph_properties{
			average_degree: 2.0,
			maximum_degree: 4,
			vertex_degree:  []int32{4, 1, 1, 2, 3, 3, 2, 1, 1, 3, 1, 2, 2},
		},
	},

	// graph-1
	{
		graph_defn: graph_definition{
			v: 3,
			e: 3,
			edges: [][2]int32{
				{0, 1},
				{0, 2},
				{1, 2},
			},
		},

		graph_props: graph_properties{
			average_degree: 2.0,
			maximum_degree: 2,
			vertex_degree:  []int32{2, 2, 2},
		},
	},
}

//
// create a graph from test_graph definition
//
func create_test_graph(test_graph *graph_definition) *Graph {
	g := New(test_graph.v)
	for _, edge := range test_graph.edges {
		v, w := edge[0], edge[1]
		g.AddEdge(v, w)
	}

	return g
}

//
// validate all properties on all graphs
//
func TestValidateGraphProperties(t *testing.T) {
	for i, graph := range graphs {
		G := create_test_graph(&graph.graph_defn)

		// average-degree validation on G
		{
			avg_degree := G.AverageDegree()
			exp_avg_degree := graph.graph_props.average_degree

			if avg_degree != exp_avg_degree {
				t.Logf("Error: test-graph: %d, expected-average-degree: %f, found-average-degree: %f\n",
					i, exp_avg_degree, avg_degree)
				t.Fail()
			}

		}

		// max-degree validation on G
		{
			max_degree := G.MaxDegree()
			exp_max_degree := graph.graph_props.maximum_degree

			if max_degree != exp_max_degree {
				t.Logf("Error: test-graph: %d, expected-max-degree: %d, found-max-degree: %d\n",
					i, exp_max_degree, max_degree)
				t.Fail()
			}
		}

		// vertex-degree validation on all vertices in G
		{
			for v := int32(0); v < G.V(); v++ {
				degree := G.Degree(v)
				exp_degree := graph.graph_props.vertex_degree[v]

				if degree != exp_degree {
					t.Logf("Error: test-graph: %d, vertex: %d, expected-degree: %d, found-degree: %d\n",
						i, v, exp_degree, degree)
					t.Fail()
				}
			}
		}
	}
}
