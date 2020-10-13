package solution

type node struct {
	key, val   int
	prev, next *node
}

type LRUCache struct {
	capacity int
	hash     map[int]*node
	head     *node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     map[int]*node{},
		head:     nil,
	}
}

func (this *LRUCache) Get(key int) int {
	resu := -1
	if temp, ok := this.hash[key]; ok {
		resu = temp.val

		if temp == this.head {
			return resu
		}
		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		temp.next = this.head
		temp.prev = this.head.prev
		this.head.prev.next = temp
		this.head.prev = temp
		this.head = temp
	}
	return resu
}

func (this *LRUCache) Put(key int, value int) {
	if temp, ok := this.hash[key]; ok {
		temp.val = value

		temp.prev.next = temp.next
		temp.next.prev = temp.prev
		temp.next = this.head
		temp.prev = this.head.prev
		this.head.prev.next = temp
		this.head.prev = temp
		this.head = temp
		return
	}
	if len(this.hash) == 0 {
		this.head = &node{
			key: key,
			val: value,
		}
		this.hash[key] = this.head
		this.head.prev = this.head
		this.head.next = this.head
		return
	}
	if len(this.hash) < this.capacity {
		temp := &node{
			key: key,
			val: value,
		}
		this.hash[key] = temp
		temp.prev = this.head.prev
		temp.next = this.head
		this.head.prev.next = temp
		this.head.prev = temp
		this.head = temp
		return
	}
	// println("->key: ", key, len(this.hash), this.capacity, this.head.prev.key, this.head.prev.prev.key)
	// need evict ...
	// println(this.head.prev.key)
	delete(this.hash, this.head.prev.key)
	this.head.prev.key = key
	this.head.prev.val = value
	this.head = this.head.prev
	this.hash[key] = this.head
}

// func (this *LRUCache) Travel() {
// 	ptr := this.head
// 	print(ptr.val, " ")
// 	ptr = ptr.next
// 	for ptr != this.head {
// 		print(ptr.val, " ")
// 		ptr = ptr.next
// 	}
// 	print(" | ")

// 	ptr = this.head.prev
// 	print(ptr.val, " ")
// 	ptr = ptr.prev
// 	for ptr != this.head.prev {
// 		print(ptr.val, " ")
// 		ptr = ptr.prev
// 	}
// 	println("...")
// }
