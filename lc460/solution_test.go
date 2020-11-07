package solution

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	// 2 dict (kv + <freq, Node>)+ dl + Node
	// I dont understand the test case , why evict 2 and then evict 4
	// why not evict 1 and 2  or 2 and 4, instead test case evict 2 and 1...
	// but I'm sure after understanding what the question wants to express, I will pass and I can fix code in 5min.
	lFUCache := Constructor(2)

	lFUCache.Put(1, 1)
	lFUCache.Put(2, 2)
	if lFUCache.Get(1) != 1 {
		t.Fail()
	} // return 1
	lFUCache.Put(3, 3) // evicts key 2
	if lFUCache.Get(2) != -1 {
		t.Fail()
	} // return -1 (not found)
	if lFUCache.Get(3) != 3 {
		t.Fail()
	} // return 3
	lFUCache.Put(4, 4) // evicts key 1.
	if lFUCache.Get(1) != -1 {
		t.Fail()
	} // return -1 (not found)
	if lFUCache.Get(3) != 3 {
		t.Fail()
	} // return 3
	if lFUCache.Get(4) != 4 {
		t.Fail()
	} // return 4

}
