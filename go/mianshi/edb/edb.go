/**
 * @desc
 * @date 2022/11/5
 * @user yangshuo
 */
//有一个二维平面RPG游戏，想要设计一个自动寻路系统，可以帮助玩家找到到达任务目标点的路线，地图是一个n * m大小的矩形。
//第一条线分别是地图的长度和宽度，然后是地图，0表示可以通过，1表示无法通过的障碍，x表示玩家位置，y表示任务坐标点.
//结果是一个数组，其中每个项目都是一个坐标，整个集合记录了玩家从起点到时钟点的路线坐标，如果玩家无法到达任务目标点，则返回空数组。


package edb

type coordinate struct {
	x int
	y int
}

//思路：回溯。1.找到起点，2.从起点开始向上下左右四个方向dfs，3.每走一个位置add到数组中，4.找到y就把路径队列放到res。
func Pathfinding(maps [][]string) [][]coordinate {
	m := len(maps)
	n := len(maps[0])

	res := [][]coordinate{}
	path := []coordinate{}
	//1.找起点
	//2.统计步数
	starti, startj := 0, 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maps[i][j] == "x" {
				startj = i
				startj = j
				continue
			}
		}
	}

	var dfs func(grid [][]string, i, j int)
	dfs = func(grid [][]string, i, j int) {
		if len(res) > 0 {
			return
		}
		//边界
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}

		//障碍
		if grid[i][j] == "1" {
			return
		}

		//走过
		if grid[i][j] == "2" {
			return
		}

		path = append(path, coordinate{i,j})
		//终点
		if grid[i][j] == "y" {
			res = append(res, path)
			path = path[0:0]
			return
		}

		//走过标记
		grid[i][j] = "2"

		//right
		dfs(grid, i, j+1)
		//down
		dfs(grid, i+1, j)
		//up
		dfs(grid, i-1, j)
		//left
		dfs(grid, i, j-1)

		grid[i][j] = "0"
	}

	dfs(maps, starti, startj)

	return res
}

//不剪枝
func PathFinding(grid [][]string) [][2]int {

	m := len(grid)
	n := len(grid[0])

	//1.找起点
	//2.统计步数
	start := [2]int{}
	count := m*n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "1" {
				count--
			}

			if grid[i][j] == "x" {
				start[0] = i
				start[1] = j
				continue
			}
		}
	}

	res := [][2]int{}
	pathes := [][2]int{}

	dirs := [][2]int{{-1,0},{0,1},{1,0},{0,-1}}//上右下左
	// 记录所有走过的节点
	vis := make(map[[2]int]bool)

	var dfs func(i, j, path int)
	dfs = func(i, j, path int) {
		pathes = append(pathes, start)
		//终点
		if grid[i][j] == "y" {
			//if len(res) == 0 {
			//	for _, v := range pathes {
			//		res = append(res, v)
			//	}
			//} else if len(res) > len(pathes) {
			//	for _, v := range pathes {
			//		res = append(res, v)
			//	}
			//}
			res = pathes
			//copy(res, pathes)
			return
		}

		//走完了还没找到
		if path >= count {
			pathes = pathes[0:0]
			return
		}

		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && !vis[[2]int{x, y}] && grid[x][y] != "1" {
				vis[[2]int{x, y}] = true
				dfs(x, y, path+1)
				delete(vis, [2]int{x, y})
			}
		}

	}

	vis[start] = true
	dfs(start[0], start[1], 1)

	return res
}