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

Performance
-----------
Machine: Intel(R) Core(TM) i7 CPU 920  @ 2.67GHz

```
- queue
    - BenchmarkPush-8    29.8    ns/op
    - BenchmarkPop-8     6.24    ns/op
- stack
    - BenchmarkPush-8    26.3    ns/op
    - BenchmarkPop-8     6.41    ns/op
- graph
- dfs
    - BenchmarkDepthFirstSearch-8        7586    ns/op
- bfs
    - BenchmarkBreadthFirstSearch-8      298089          ns/op
- traversal
    - BenchmarkBFSTraversal-8    285810          ns/op
    - BenchmarkDFSTraversal-8    11488   ns/op
- line_parser
```
