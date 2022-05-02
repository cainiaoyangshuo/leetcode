/**
 * @desc
 * @date 2022/4/29
 * @user yangshuo
 */
package _05hashSet

import "container/list"

const base = 779

type MyHashSet struct {
	data []list.List
}


func Constructor() MyHashSet {
	return MyHashSet{
		data: make([]list.List, base),
	}
}


func (this *MyHashSet) Add(key int)  {
	if !this.Contains(key) {
		idx := this.hash(key)
		this.data[idx].PushBack(key)
	}
}


func (this *MyHashSet) Remove(key int)  {
	idx := this.hash(key)

	for e := this.data[idx].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			this.data[idx].Remove(e)
		}
	}
}


func (this *MyHashSet) Contains(key int) bool {
	idx := this.hash(key)

	for e := this.data[idx].Front(); e != nil; e = e.Next() {
		if e.Value == key {
			return true
		}
	}

	return false
}

func (this *MyHashSet) hash(key int) int {
	return key % base
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
