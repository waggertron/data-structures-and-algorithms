package graphs

type Graph struct {
	Nodes map[string][]string
}

func MakeGraph(edges [][]string) map[string][]string {
	g := make(map[string][]string)

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]

		if _, ok := g[n1]; !ok {
			g[n1] = []string{n2}
		}
	}

	return g
}


