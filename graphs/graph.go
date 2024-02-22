package graphs

type Graph struct {
	Nodes map[int][][2]int
}

func MakeGraph(edges [][3]int) map[int][][2]int {
	nodes := make(map[int][][2]int)

	for _, edge := range edges {
		src, dest, weight := edge[0], edge[1], edge[2]

		_, ok := nodes[src]
		if !ok {
			nodes[src]  = make([][2]int, 1)
		}
		edges := nodes[src]
		nodes[src] = append(edges, [2]int{dest, weight})
	}

	return nodes
}
