package main

import "fmt"

func dfs(image [][]int, i, j, c0, c1 int) {
	if i < 0 || j < 0 || i >= len(image) || j >= len(image[0]) || image[i][j] != c0 {
		return
	}
	image[i][j] = c1
	dfs(image, i, j-1, c0, c1)
	dfs(image, i, j+1, c0, c1)
	dfs(image, i-1, j, c0, c1)
	dfs(image, i+1, j, c0, c1)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, image[sr][sc], newColor)
	}
	return image
}

func main() {
	fmt.Println(floodFill([][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2))
}
