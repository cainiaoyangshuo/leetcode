/**
 * @desc 933. 最近的请求次数 https://leetcode-cn.com/problems/number-of-recent-calls/
 * 思路：队列，将时间t入队，不断的弹出队首早于t-3000的时间，循环结束后队列长度即为t-3000时间内请求个数。
 * @date 2022/5/6
 * @user yangshuo
 */
package _33_RecentCounter

type RecentCounter struct {
	q []int
}


func Constructor() RecentCounter {
	return RecentCounter {
		q: []int{0},
	}
}


func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.q = this.q[1:]
	}

	return len(this.q)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
