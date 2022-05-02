/**
 * @desc
 * @date 2022/4/30
 * @user yangshuo
 */
package _06hashMap

import "container/list"

const base = 779

type entry struct {
	key, value int
}

type MyHashMap struct {
	data []list.List
}


func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]list.List, base),
	}
}


func (this *MyHashMap) Put(key int, value int)  {
	idx := this.hash(key)

	for e := this.data[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[idx].PushBack(entry{key: key, value: value})
}


func (this *MyHashMap) Get(key int) int {
	idx := this.hash(key)

	for e := this.data[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}

	return -1
}


func (this *MyHashMap) Remove(key int)  {
	idx := this.hash(key)

	for e := this.data[idx].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			this.data[idx].Remove(e)
		}
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % base
}
/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
