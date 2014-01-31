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
// this package implements the symbol graph where vertex names are
// strings and number of edges/vertices are implicitly defined. this
// is more typical of real-world (tm) graph applications
//
package symbol_graph

import (
	"bufio"
	"fmt"
	"github.com/anupamk/common-utilz/graph"
	"github.com/anupamk/common-utilz/line_parser"
	"io"
	"os"
)

type SymbolGraph struct {
	sym_table map[string]int32 // vertex-name -> vertex-index
	keys      []string         // vertex-index -> vertex-name
	sym_graph *graph.Graph     // resultant graph
}

type string_slice_t []string

//
// this function returns the stringified representation of the
// symbol-graph
//
func (sg *SymbolGraph) String() (retval string) {

	retval += fmt.Sprintf("symbol-table: %v\n", sg.sym_table)
	retval += fmt.Sprintf("reverse-index: %v\n", sg.keys)
	retval += fmt.Sprintf("graph\n%s", sg.sym_graph)

	return
}

//
// this function is called to create a symbol-graph from it's
// serialized definition.
//
func LoadFromReader(src *bufio.Reader, sep string) (sg *SymbolGraph, sg_err error) {
	var edge_list []string_slice_t

	sg = &SymbolGraph{
		sym_table: make(map[string]int32),
		keys:      nil,
		sym_graph: nil,
	}

	// populate the symbol table and
	for vlist, err := line_parser.StringsFromReader(src, '#', sep); err != io.EOF; {
		edge_list = append(edge_list, vlist)

		for _, v := range vlist {
			if _, ok := sg.sym_table[v]; ok {
				continue
			}
			sg.sym_table[v] = int32(len(sg.sym_table))
		}

		// next-line
		vlist, err = line_parser.StringsFromReader(src, '#', sep)
	}

	// the reverse index
	sg.keys = make([]string, len(sg.sym_table))
	for sk, sv := range sg.sym_table {
		sg.keys[sv] = sk
	}

	// create the symbol graph
	sg.sym_graph = graph.New(int32(len(sg.sym_table)))
	for _, vlist := range edge_list {
		sv_id := sg.sym_table[vlist[0]]
		for i := 1; i < len(vlist); i++ {
			dv_id := sg.sym_table[vlist[i]]
			sg.sym_graph.AddEdge(sv_id, dv_id)
		}
	}

	return
}

//
// this is a convenience interface over LoadSymbolGraphFromReader(...)
// to create a symbol-graph from its serialized definition stored in a
// file identified by 'fname'
//
func LoadFromFile(fname string, sep string) (sg *SymbolGraph, err error) {
	var f *os.File

	if f, err = os.Open(fname); err != nil {
		return
	}
	defer f.Close()

	file_reader := bufio.NewReader(f)
	if sg, err = LoadFromReader(file_reader, sep); err != nil {
		return
	}

	return
}

//
// this function returns true if the symbol-graph contains 'str' as a
// vertex
//
func (sg *SymbolGraph) Contains(str string) bool {
	_, ok := sg.sym_table[str]
	return ok
}

//
// this function returns the index associated with a given key, if the
// key doesn't exist, it panics. thus clients, are expected to ensure
// that the key is available before coming here.
//
func (sg *SymbolGraph) Index(key string) (idx int32) {
	var ok bool

	if idx, ok = sg.sym_table[key]; !ok {
		err := fmt.Errorf("key: '%s' doesn't exist", key)
		panic(err)
	}

	return
}

//
// this function returns the name associated with a given vertex. if
// the vertex is invalid, an error is flagged
func (sg *SymbolGraph) Name(vertex_id int32) (name string, err error) {
	if vertex_id > int32(len(sg.keys)) {
		err = fmt.Errorf("vertex: '%d' doesn't exist", vertex_id)
		return
	}

	name = sg.keys[vertex_id]
	return
}

//
// this function returns a pointer to the underlying graph in the
// symbol graph.
func (sg *SymbolGraph) G() (g *graph.Graph) { return sg.sym_graph }
