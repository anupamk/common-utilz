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
// this file implements the shortest path from source to destination
// vertices in symbol graphs
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/anupamk/common-utilz/graph/symbol_graph"
	"github.com/anupamk/common-utilz/graph/traversal"
	"io"
	"log"
	"os"
	"strings"
)

//
// program input control
//
var (
	data_file     string // data-file
	separator     string // field-separator
	source        string // source-vertex
	verbose_debug bool   // just-what-it-sez
)

// setup various command line parameters
func init() {
	// application parameterz
	flag.StringVar(&data_file, "input-file", "", "symbol-graph data file name")
	flag.StringVar(&separator, "separator", "", "field separator in the input")
	flag.StringVar(&source, "source", "", "source vertex")

	// debugging stuff
	flag.BoolVar(&verbose_debug, "debug", true, "generate verbose debugging")

}

// rudimentary checks and parameter cleanups
func sanitize_and_validate_cmdline_params() {
	if data_file = strings.TrimSpace(data_file); len(data_file) == 0 {
		fmt.Fprintf(os.Stderr, "usage-error: bad data-file '%s'\n", data_file)
		flag.Usage()

		os.Exit(255)
	}

	// don't mess with the separator
	if len(separator) == 0 {
		fmt.Fprintf(os.Stderr, "usage-error: bad seperator '%s'\n", separator)
		flag.Usage()

		os.Exit(254)
	}

	if source = strings.TrimSpace(source); len(source) == 0 {
		fmt.Fprintf(os.Stderr, "usage-error: bad source '%s'\n", source)
		flag.Usage()

		os.Exit(253)
	}
}

func main() {
	flag.Parse()
	sanitize_and_validate_cmdline_params()

	if verbose_debug {
		log.Printf("-- degree-of-seperation parameter dump: input-file: '%s', separator: '%s', source-vertex: '%s' --\n",
			data_file, separator, source)
	}

	// parse the file-name and create a symbol graph
	sg, err := symbol_graph.LoadFromFile(data_file, separator)
	if err != nil {
		log.Fatal(err)
	}

	if verbose_debug {
		sg_g := sg.G()
		log.Printf("-- symbol-graph created. vertices: '%d', edges: '%d' --\n", sg_g.V(), sg_g.E())
	}

	if !sg.Contains(source) {
		fmt.Fprintf(os.Stderr, "fatal-error: '%s' not found in data-base\n", source)
		flag.Usage()

		os.Exit(252)
	}

	//
	// ok, so by now, we have a valid graph, and a valid source,
	// let's answer some questions...
	//

	// determine all paths from source -> other vertices on the
	// graph
	all_paths, err := traversal.SingleSourceShortestPaths(sg.G(), sg.Index(source))
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal-error: invalid source '%s'\n", source)
		os.Exit(128)
	}

	// query-rsp loop
	for stdin_reader := bufio.NewReader(os.Stdin); ; {
		fmt.Fprintf(os.Stdout, "degree-of-separation --> ")

		// read and sanitize input
		dest_name, err := stdin_reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		if dest_name = strings.TrimSpace(dest_name); len(dest_name) == 0 {
			continue
		}

		if verbose_debug {
			log.Printf("-- user-input: '%s' --\n", dest_name)
		}

		// got something useful, let's see do we know it ?
		if !sg.Contains(dest_name) {
			log.Printf("error: '%s' doesn't exist\n", err)
			continue
		}

		// is there a path to destination ?
		dst_vertex := sg.Index(dest_name)

		if !all_paths.PathExists(dst_vertex) {
			log.Printf("no path to: '%s'\n", dest_name)
			continue
		}

		// yes there is a path, enumerate it...
		src_dst_path := all_paths.PathTo(dst_vertex)

		fmt.Fprintf(os.Stdout, "%s\n", source)
		for _, v := range src_dst_path[1:] {
			vertex_name, _ := sg.Name(v)
			fmt.Fprintf(os.Stdout, "    %s\n", vertex_name)
		}
	}
}
