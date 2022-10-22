/**
 * @desc 733.图像渲染 https://leetcode.cn/problems/flood-fill/
 * @date 2022/5/30
 * @user yangshuo
 */

package _33_floodFill
/**
 * 广度优先搜索
 * 从目标位置开始上色，若像素颜色和color相同，则改变颜色为newcolor；然后再从四个方向进行上色，重复上述过程。
 */
// 创建一个结构体，存储当前结点的x,y值

type QueueNode struct {
	x int
	y int
}

// bfs 广度优先搜索。 通过queue
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// 创建一个数组
	queue := make([]QueueNode, 0)
	// 把初始结点的元素放到队列中(入队)
	queue = append(queue, QueueNode{sr, sc})
	// 存一下初始结点的值
	firstColor := image[sr][sc]
	// 如果我们要改变的值和我们的新值是一样的，直接返回，因为不需要改变。 如果没有这个判断，会死循环。
	if firstColor == newColor {
		return image
	}
	// 队列为空的时候遍历结束。
	for len(queue) != 0 {
		// 获取队列第一个元素（出队）
		value := queue[0]
		queue = queue[1:]
		// 把newColor给当前坐标的值
		image[value.x][value.y] = newColor
		// 然后比对当前坐标四周的结点，有没有符合条件的，如果有，添加到队列中。 循环这个过程。
		if value.x-1 >= 0 && image[value.x-1][value.y] == firstColor {
			queue = append(queue, QueueNode{value.x - 1, value.y})
		}
		if value.x+1 < len(image) && image[value.x+1][value.y] == firstColor {
			queue = append(queue, QueueNode{value.x + 1, value.y})
		}
		if value.y-1 >= 0 && image[value.x][value.y-1] == firstColor {
			queue = append(queue, QueueNode{value.x, value.y - 1})
		}
		if value.y+1 < len(image[value.x]) && image[value.x][value.y+1] == firstColor {
			queue = append(queue, QueueNode{value.x, value.y + 1})
		}
	}
	return image
}