package _46LRU

import (
	"fmt"
	"testing"
)

func Test_LRU(t *testing.T)  {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1) // 缓存是 {1=1}
	lRUCache.Put(1, 1) // 缓存是 {1=1, 2=2}
	//lRUCache.Get(1)   // 返回 1
	lRUCache.Put(2, 3) // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//lRUCache.Get(4)    // 返回 -1 (未找到)
	lRUCache.Put(4, 1) // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	fmt.Printf("%d\n ", lRUCache.Get(1) ) // 返回 -1 (未找到)
	fmt.Printf("%d\n ", lRUCache.Get(2) )   // 返回 3
	//lRUCache.Get(4)    // 返回 4
}
