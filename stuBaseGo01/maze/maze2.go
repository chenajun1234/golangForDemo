package main

import (
	"fmt"
	"os"
)

type points struct {
	r, c int
}

var dirArray = [4]points{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func readMaze2(fileName string) [][]int {
	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("os open err", e)
		return nil
	}
	defer file.Close()
	var row, col int
	_, err := fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println(row, col)
	if err != nil {
		fmt.Println("os open err", err)
		return nil
	}
	maze := make([][]int, row)
	for r := range maze {
		maze[r] = make([]int, col)
		for c := range maze[r] {
			fmt.Fscanf(file, "%d", &maze[r][c])
		}
	}
	return maze
}

func (p *points) addP(dir points) points {
	return points{p.r + dir.r, p.c + dir.c}
}

func (p *points) at(maze [][]int) (int, bool) {
	if p.r < 0 || p.r >= len(maze) {
		return 0, false
	}
	if p.c < 0 || p.c >= len(maze[0]) {
		return 0, false
	}
	return maze[p.r][p.c], true
}
func walk2(maze [][]int, strat points, end points) [][]int {
	//初始化步骤
	steps := make([][]int, len(maze))
	for r := range steps {
		steps[r] = make([]int, len(maze[r]))
	}
	//初始化能够探索的队列
	Q := []points{strat}
	//循环队列
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		//探索下一个
		for _, dir := range dirArray {
			next := cur.addP(dir)

			if next == strat {
				continue
			}
			val, ok := next.at(maze)
			if !ok || val != 0 {
				continue
			}
			val, ok = next.at(steps)
			fmt.Println("at val", val, ok)
			if !ok || val != 0 {
				continue
			}
			fmt.Println("steps val", val, ok)
			curVal, _ := cur.at(steps)
			steps[next.r][next.c] = 1 + curVal
			Q = append(Q, next)
		}

	}
	return steps
}
func main() {
	var fileName = "F:/work/gostu/src/stuBaseGo01/maze/maze.in"
	maze := readMaze2(fileName)
	//fmt.Println(maze)
	steps := walk2(maze, points{0, 0}, points{len(maze) - 1, len(maze[0]) - 1})
	fmt.Println(steps)
	for r := range steps {
		for c := range steps[r] {
			fmt.Printf("%3d", steps[r][c])
		}
		fmt.Println()
	}

}
