package dynamicprogramming

import "fmt"

func UniquePaths2(obstacleGrid [][]int) int {
	y := len(obstacleGrid)
	if y == 0 {
		return 0
	}
	x := len(obstacleGrid[0])
	if x == 0 {
		return 0
	}

	table := make([][]int, y)

	for i := range table {
		table[i] = make([]int, x)
	}

	table[0][0] = 0

	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
		}
	}

	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {

		}
	}

	prettyPrint(table)
	prettyPrint(obstacleGrid)

	return table[y][x]

}

func prettyPrint(t [][]int) {
	for _, row := range t {
		fmt.Println(row)
	}
}
