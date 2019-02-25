package main

import (
	"fmt"
	"os"
)

/**
广度优先搜索 算法
*/
type point struct {
	r, c int
}

/**
读取迷宫数据
*/
func readMaze(fileName string) [][]int {
	file, e := os.Open(fileName)
	if e != nil {
		fmt.Println("os open err", e)
	}
	defer file.Close()
	var row, col int
	n, err := fmt.Fscanf(file, "%d %d", &row, &col) //一次读两个数字
	if err != nil {
		fmt.Println("Fscanferr", err)
	}
	fmt.Println(row, col, n)

	maze := make([][]int, row)
	for r := range maze {
		maze[r] = make([]int, col)
		for c := range maze[r] {
			fmt.Fscanf(file, "%d", &maze[r][c])
		}
	}
	return maze
}

//探索方向
var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

//探索下一个
func (p *point) add(dir point) point {
	return point{p.r + dir.r, p.c + dir.c}
}

//坐标点 迷宫的值
func (p *point) at(maze [][]int) (int, bool) {
	if p.r < 0 || p.r >= len(maze) {
		return 0, false
	}
	if p.c < 0 || p.c >= len(maze[0]) {
		return 0, false
	}
	return maze[p.r][p.c], true
}

func walk(maze [][]int, strat point, end point) [][]int {
	//步数
	steps := make([][]int, len(maze))
	for r := range maze {
		steps[r] = make([]int, len(maze[r]))
	}
	//能探索的队列
	q := []point{strat}
	//探索队列里面的点
	for len(q) > 0 {
		curr := q[0]
		q = q[1:] //探索完毕 去掉
		if curr == end {
			break
		}
		//探索点的 上左下游
		for _, dir := range dirs {
			next := curr.add(dir)
			if next == strat {
				continue
			}
			//获取next 位置
			val, ok := next.at(maze)
			if !ok || val != 0 { //在迷宫已经探索了
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 { //前面已经探索了
				continue
			}
			//加入步骤
			currVal, _ := curr.at(steps)
			steps[next.r][next.c] = currVal + 1
			q = append(q, next)
		}
	}
	return steps
}
func main() {
	var fileName = "F:/work/gostu/src/stuBaseGo01/maze/maze.in"
	maze := readMaze(fileName)
	/*for r := range maze {
		for c := range maze[r] {
			fmt.Print(maze[r][c])
		}
		fmt.Println()
	}*/
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for r := range steps {
		for c := range steps[r] {
			fmt.Printf("%3d", steps[r][c])
		}
		fmt.Println()
	}
}
