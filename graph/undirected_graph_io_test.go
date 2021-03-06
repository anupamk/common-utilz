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
	"bufio"
	"fmt"
	"strings"
	"testing"
)

var graph_data_files = [...]string{
	"data/graph-001.data",
	"data/graph-002.data",
}

//
// save a newly created graph, and see if the serialized-output
// matches the expected one
//
func Example_001_CreateAndDumpGraph() {
	tmp, _ := LoadFromFile(graph_data_files[0])
	tmp_str := tmp.Serialize()
	fmt.Println(tmp_str)

	// Output:
	// 13
	// 13
	// 0 6
	// 0 2
	// 0 1
	// 0 5
	// 3 5
	// 3 4
	// 4 5
	// 4 6
	// 7 8
	// 9 11
	// 9 10
	// 9 12
	// 11 12
}

func Example_002_CreateAndDumpGraph() {
	tmp, _ := LoadFromFile(graph_data_files[1])
	tmp_str := tmp.Serialize()
	fmt.Println(tmp_str)

	// Output:
	// 3
	// 3
	// 0 2
	// 0 1
	// 1 2
}

//
// create a new graph from serialized graph, and compare the two for
// equality.
//
func TestLoadGraph(t *testing.T) {
	for _, fname := range graph_data_files {
		g1, _ := LoadFromFile(fname)
		g1_str := g1.Serialize()

		g1_reader := strings.NewReader(g1_str)
		graph_reader := bufio.NewReader(g1_reader)

		g2, _ := LoadFromReader(graph_reader)

		if cmp_graph(g1, g2) == false {
			t.Log("Error: Unequal Graphs")
			t.Logf("original-graph:\n%s", g1)
			t.Logf("new-graph:\n%s", g2)
			t.Fail()
		}
	}
}
