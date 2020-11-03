package solution

// use a bucket replace plain Node to get a O(1) element move
// when exists many elements with same value
// hashmap + bucket + double linked list 的思路
// 逻辑太长，难以debug
// 要判断的边界很多
// 显然这种方法只适合工作中使用，并不适合10min内解题
type Bucket struct {
	Val        int
	Next, Prev *Bucket
	Children   *Node
}

type Node struct {
	Key        string
	Next, Prev *Node
	Bkt        *Bucket
}

type AllOne struct {
	// head = min
	head, tail *Bucket
	hash       map[string]*Node
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	return AllOne{
		head: nil,
		tail: nil,
		hash: map[string]*Node{},
	}
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	var ptr *Bucket
	var ptr2 *Node
	ptr = this.head
	for ptr != nil {
		ptr2 = ptr.Children
		print(ptr.Val)
		for ptr2 != nil {
			print(" ", ptr2.Key)
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
		println()
	}
	if v, ok := this.hash[key]; ok {
		// if 上个bkt刚好+1，移到上个桶
		// remove v to prev bucket
		if v.Bkt.Prev != nil && v.Bkt.Val+1 == v.Bkt.Prev.Val {
			v.Bkt = v.Bkt.Prev
			// v.Prev.Next = v.Next
			// v.Next.Prev = v.Prev
			v.Prev = nil
			v.Next = v.Bkt.Children
			v.Bkt.Children = v
			v.Next.Prev = v
		} else {
			newBkt := &Bucket{
				Val:  v.Bkt.Val + 1,
				Prev: v.Bkt.Prev,
				Next: v.Bkt,
			}
			v.Bkt = newBkt
			if v.Bkt.Prev != nil {
				v.Bkt.Prev.Next = newBkt
				v.Bkt.Prev = newBkt
			} else {
				v.Bkt.Prev = newBkt
				this.head = newBkt
			}
		}
	} else {
		this.hash[key] = &Node{
			Key: key,
		}
		if this.tail != nil && this.tail.Val == 1 {
			this.hash[key].Bkt = this.tail
			this.hash[key].Next = this.tail.Children
			this.tail.Children.Prev = this.hash[key]
			this.tail.Children = this.hash[key]
		} else {
			newTailBkt := &Bucket{
				Val:      1,
				Children: this.hash[key],
			}
			this.hash[key].Bkt = newTailBkt
			if this.tail == nil {
				this.head = newTailBkt
				this.tail = newTailBkt
			} else {
				this.tail.Next = newTailBkt
				newTailBkt.Prev = this.tail
				this.tail = newTailBkt
			}
		}
	}

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {

}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail == nil {
		return ""
	}
	return this.tail.Children.Key
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head == nil {
		return ""
	}
	return this.head.Children.Key
}
