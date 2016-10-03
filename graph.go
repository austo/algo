package main

type edgeNode struct {
	y      int
	weight int
	next   *edgeNode
}

type Graph struct {
	edges    []*edgeNode /* adjency info */
	degree   []int       /* outdegree of each vertex */
	nedges   int         /* number of edges in graph */
	directed bool        /* is graph directed? */
}

func New(vertices int) *Graph {
	return &Graph{
		edges:  make([]*edgeNode, vertices),
		degree: make([]int, vertices),
	}
}

func (g *Graph) AddEdge(x, y int, directed bool) {
	p := &edgeNode{
		y:    y,
		next: g.edges[x],
	}

	g.edges[x] = p
	g.degree[x]++

	if !directed {
		g.AddEdge(y, x, true)
	} else {
		g.edges++
	}
}

func (g *Graph) Bfs()
