common-utilz
============
Provides an implementation of a small set of commonly used
data-structures, algorithms and some conveinence routines for
day-to-day programming activities.

List Of Available Stuff
-----------------------
Convenience Routines:
  - [Line Parser](http://godoc.org/github.com/anupamk/common-utilz/line_parser)

Data Structures:
  - [Queue](http://godoc.org/github.com/anupamk/common-utilz/queue)
  - [Graph](http://godoc.org/github.com/anupamk/common-utilz/graph)
  - [Stack](http://godoc.org/github.com/anupamk/common-utilz/stack)

Algorithms:
  - Graphs
      - [Depth First Search](http://godoc.org/github.com/anupamk/common-utilz/graph/dfs)
      - [Breadth First Search](http://godoc.org/github.com/anupamk/common-utilz/graph/bfs)

Performance
-----------
Machine: Intel(R) Core(TM) i7 CPU 920  @ 2.67GHz

```
- queue
    - BenchmarkPush-8    29.8    ns/op
    - BenchmarkPop-8     6.24    ns/op
- graph
- dfs
    - BenchmarkDepthFirstSearch-8        7586    ns/op
- line_parser
```
