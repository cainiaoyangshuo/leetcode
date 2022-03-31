package offer076_topK

import "container/heap"

type Heap []int

func (h *Heap) Push(v interface{})  {
	*h = append(*h, v.(int))
}

func (h *Heap) Pop() interface{} {
	tmp := *h
	v := tmp[len(tmp)-1]
	*h = tmp[0:len(tmp)-1]
	return v
}

func (h Heap) Len() int {
	return len(h)
}

func (h *Heap) Top() int {
	return (*h)[0]
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

// Swap swaps the elements with indexes i and j.
func (h Heap) Swap(i, j int) {
	h[i], h[j] =  h[j],  h[i]
}

func findKthLargest(nums []int, k int) int {
	length := len(nums)

	h := &Heap{}
	heap.Init(h)
	for i:=0;i<length;i++ {
		val := nums[i]
		if h.Len() < k  {
			h.Push(val)
		} else {
			top := h.Top()
			if top < val {
				h.Pop()
				h.Push(val)
			}
		}
	}

	return h.Top()
}
