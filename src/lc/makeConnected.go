package lc

func MakeConnected(n int, connections [][]int) (ans int) {
		if len(connections) < n-1 {
			return -1
		}

		graph := make([][]int, n)
		for _, c := range connections {
			v, w := c[0], c[1]
			graph[v] = append(graph[v], w)
			graph[w] = append(graph[w], v)
		}

		vis := make([]bool, n)
		var dfs func(int)
		dfs = func(from int) {
			vis[from] = true
			for _, to := range graph[from] {
				if !vis[to] {
					dfs(to)
				}
			}
		}
		for i, v := range vis {
			if !v {
				ans++
				dfs(i)
			}
		}
		return ans - 1
}
