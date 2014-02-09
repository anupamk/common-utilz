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
//
// this package implements the symbol digraph where vertex names are
// strings and number of edges/vertices are implicitly defined. this
// is more typical of real-world (tm) graph applications
//
// this file implements the testing routine for symbol-digraphs
//
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

	for _, v := range topo_order {
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
