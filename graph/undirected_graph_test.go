//
// Copyright (c) 2014, Anupam Kapoor. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// Author: anupam.kapoor@gmail.com (Anupam Kapoor)
//
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
