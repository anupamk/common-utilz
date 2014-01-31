common-utilz
============
Provides an implementation of a small set of commonly used
data-structures, algorithms and some conveinence routines for
day-to-day programming activities.

List Of Available Stuff
-----------------------
Convenience Routines:
  - [Line Parser](http://godoc.org/github.com/anupamk/common-utilz/line_parser)
  - [Slice Utils](http://godoc.org/github.com/anupamk/common-utilz/slice_utils)

Data Structures:
  - [Queue](http://godoc.org/github.com/anupamk/common-utilz/queue)
  - [Graph](http://godoc.org/github.com/anupamk/common-utilz/graph)
  - [Stack](http://godoc.org/github.com/anupamk/common-utilz/stack)

Algorithms:
  - Graphs
      - [Depth First Search](http://godoc.org/github.com/anupamk/common-utilz/graph/dfs)
      - [Breadth First Search](http://godoc.org/github.com/anupamk/common-utilz/graph/bfs)
      - [Traversal](http://godoc.org/github.com/anupamk/common-utilz/graph/traversal)
      - [Symbol Graph](http://godoc.org/github.com/anupamk/common-utilz/graph/symbol_graph)
      - [Misc](http://godoc.org/github.com/anupamk/common-utilz/graph/algorithms)

Performance
-----------
Machine: Intel(R) Core(TM) i7 CPU       Q 720  @ 1.60GHz

```
- line_parser
- graph
- symbol_graph
- bfs
    - BenchmarkBreadthFirstSearch-8      414799  ns/op
- applications
- dfs
    - BenchmarkDepthFirstSearch-8        10655   ns/op
- traversal
    - BenchmarkBFSGraphTraversal-8       677     ns/op
    - BenchmarkDFSGraphTraversal-8       647     ns/op
- algorithms
    - BenchmarkConnectedComponents-8     0.76    ns/op
    - BenchmarkBFSPathTo-8       	 394695  ns/op
    - BenchmarkDFSPathTo-8       	 390624  ns/op
- queue
    - BenchmarkPush-8    		 54.5    ns/op
    - BenchmarkPop-8     		 7.27    ns/op
- stack
    - BenchmarkPush-8    		 52.5    ns/op
    - BenchmarkPop-8     		 7.42    ns/op
- slice_utils
```
