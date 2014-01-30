//
// this file implements a client which loads a symbol graph from a
// specified file, and then processes clients query for adjacency
// information for a given vertex
//
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/anupamk/common-utilz/graph/symbol_graph"
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
	separator     string // field-sepe
	verbose_debug bool   // just-what-it-sez
)

// setup various command line parameters
func init() {
	// application parameterz
	flag.StringVar(&data_file, "input-file", "", "symbol-graph data file name")
	flag.StringVar(&separator, "separator", "", "field separator in the input")

	// debugging stuff
	flag.BoolVar(&verbose_debug, "debug", true, "generate verbose debugging")

}

// rudimentary checks
func validate_cmdline_params() {
	if len(data_file) == 0 {
		fmt.Fprintf(os.Stderr, "usage-error: bad data-file '%s'\n", data_file)
		flag.Usage()

		os.Exit(255)
	}

	if len(separator) == 0 {
		fmt.Fprintf(os.Stderr, "usage-error: bad seperator '%s'\n", separator)
		flag.Usage()

		os.Exit(254)
	}
}

func main() {
	flag.Parse()
	validate_cmdline_params()

	if verbose_debug {
		log.Printf("-- creating symbol graph, input-file: '%s', separator: '%s' --\n", data_file, separator)
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

	// handle user queries from stdin
	for sg_g, stdin_reader := sg.G(), bufio.NewReader(os.Stdin); ; {
		// our prompt
		fmt.Fprintf(os.Stdout, "graph-client --> ")

		// read input
		line_in, err := stdin_reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}

		if line_in = strings.TrimSpace(line_in); len(line_in) == 0 {
			continue
		}

		if verbose_debug {
			log.Printf("-- user-input: '%s' --\n", line_in)
		}

		// got something useful, let's see do we know it ?
		source, err := sg.Index(line_in)
		if err != nil {
			log.Printf("error: '%s'", err)
			continue
		}

		// ok we do, dump named adjacency list
		fmt.Printf("%s\n", line_in)
		adj_list := sg_g.Adj(source)
		for _, v := range adj_list {
			vname, _ := sg.Name(v)
			fmt.Printf("  %s\n", vname)
		}
	}
}
