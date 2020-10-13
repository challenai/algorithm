package solution

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	lRUCache := Constructor(2)

	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	if lRUCache.Get(1) != 1 {
		t.Fail()
	} // return 1
	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	if lRUCache.Get(2) != -1 {
		t.Fail()
	} // returns -1 (not found)
	lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	if lRUCache.Get(1) != -1 {
		t.Fail()
	} // return -1 (not found)
	// println(lRUCache.Get(3))
	if lRUCache.Get(3) != 3 {
		t.Fail()
	} // return 3
	if lRUCache.Get(4) != 4 {
		t.Fail()
	} // return 4
}
