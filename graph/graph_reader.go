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
//
// commonly used functions for reading/writing graph structures
//
package graph

import (
	"bufio"
	"fmt"
	"github.com/anupamk/common-utilz/line_parser"
	"io"
	"os"
)

//
// this function is called to create a graph from it's serialized
// definition.
//
func LoadFromReader(src *bufio.Reader) (new_graph *Graph, err error) {
	var V, E int32
	var edges [][2]int32

	V, E, edges, err = parse_graph_datafile(src)
	if err != nil && err != io.EOF {
		goto all_done
	}

	// create the graph, and setup the connections
	new_graph = New(V)
	for i := int32(0); i < E; i++ {
		v, w := edges[i][0], edges[i][1]

		// skip edges which are obviouzly bogus
		if V < v || v < 0 || V < w || w < 0 {
			fmt.Printf("skipping bogus connection: %d %d\n", v, w)
			continue
		}
		new_graph.AddEdge(w, v)
	}

all_done:
	return
}

//
// this is a convenience interface over LoadGraph(...) to create a
// graph from its serialized definition stored in a file identified by
// 'fname'
//
func LoadFromFile(fname string) (g *Graph, err error) {
	var f *os.File

	if f, err = os.Open(fname); err != nil {
		return nil, err
	}
	defer f.Close()

	file_reader := bufio.NewReader(f)
	g, err = LoadFromReader(file_reader)

	return
}

//
// this function is called to create a graph from it's serialized
// definition.
//
func LoadDigraphFromReader(src *bufio.Reader) (new_graph *Digraph, err error) {
	var V, E int32
	var edges [][2]int32

	V, E, edges, err = parse_graph_datafile(src)
	if err != nil && err != io.EOF {
		goto all_done
	}

	// create the graph, and setup the connections
	new_graph = CreateDigraph(V)
	for i := int32(0); i < E; i++ {
		v, w := edges[i][0], edges[i][1]

		// skip edges which are obviouzly bogus
		if V < v || v < 0 || V < w || w < 0 {
			fmt.Printf("skipping bogus connection: %d %d\n", v, w)
			continue
		}
		new_graph.AddEdge(w, v)
	}

all_done:
	return
}

//
// this is a convenience interface over LoadGraph(...) to create a
// graph from its serialized definition stored in a file identified by
// 'fname'
//
func LoadDigraphFromFile(fname string) (g *Digraph, err error) {
	var f *os.File

	if f, err = os.Open(fname); err != nil {
		return nil, err
	}
	defer f.Close()

	file_reader := bufio.NewReader(f)
	g, err = LoadDigraphFromReader(file_reader)

	return
}

/*
 * unexported stuff
**/

//
// parse the input data file which is expected to be in the following
// format:
//     <line-001> number-of-vertices (V)
//     <line-002> number-of-edges (E)
//     <line-003> vertex-i vertex-j (edges[][2])
//     ....................
//     ....................
//     <line-N>   vertex-i vertex-j (edges[][2])
//
// commented-lines (starting with '#') / non-empty lines are ignored
//
func parse_graph_datafile(in *bufio.Reader) (V int32, E int32, Edges [][2]int32, err error) {
	var values []int32
	var comment_char byte = '#'

	if in == nil {
		return
	}

	// parse vertex-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	V = values[0]

	// parse edge-count
	values, err = line_parser.Int32sFromReader(in, comment_char)
	if err != nil {
		return
	}
	E = values[0]

	// parse edge-list
	Edges = make([][2]int32, E)
	for i := int32(0); i < E; i++ {
		values, err = line_parser.Int32sFromReader(in, comment_char)
		if err != nil {
			break
		}
		Edges[i][0], Edges[i][1] = values[0], values[1]
	}

	return
}
