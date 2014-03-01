package symbol_graph

import (
	"bufio"
	"github.com/anupamk/common-utilz/line_parser"
	"io"
)

//
// this function creates the symbol table, and the reverse-index
// (string->int32) from a given input source and a user-defined
// seperator string.
//
// normally this would be invoked prior to acutally populating the
// appropriate symbol-graph type...
func load_symtab_revindex_from_reader(src *bufio.Reader, sep string) (symtab map[string]int32, revindex []string, edge_list []string_slice_t) {
	symtab = make(map[string]int32)

	// populate the symbol table, ignoring all comment lines
	// i.e. lines begining with '#'
	for vlist, err := line_parser.StringsFromReader(src, '#', sep); err != io.EOF; {
		edge_list = append(edge_list, vlist)

		for _, v := range vlist {
			if _, ok := symtab[v]; ok {
				continue
			}
			symtab[v] = int32(len(symtab))
		}

		// next-line
		vlist, err = line_parser.StringsFromReader(src, '#', sep)
	}

	// the reverse index
	revindex = make([]string, len(symtab))
	for sk, sv := range symtab {
		revindex[sv] = sk
	}

	return
}
