package solution

// use a bucket replace plain Node to get a O(1) element move
// when exists many elements with same value
// hashmap + bucket + double linked list 的思路
// 逻辑太长，难以debug
// 要判断的边界很多
// 显然这种方法只适合工作中使用，并不适合10min内解题
// 在看了标准解后发现也没有更好的解了。。。
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
	counter := 0
	for ptr != nil && counter < 3 {
		ptr2 = ptr.Children
		print(ptr.Val)
		counter++
		counter2 := 0
		for ptr2 != nil && counter2 < 3 {
			print(" ", ptr2.Key)
			ptr2 = ptr2.Next
			counter2++
		}
		ptr = ptr.Next
		println()
	}
	println()
	println()
	println()
	if v, ok := this.hash[key]; ok {
		// if 上个bkt刚好+1，移到上个桶
		// remove v to prev bucket
		if v.Bkt.Prev != nil && v.Bkt.Val+1 == v.Bkt.Prev.Val {
			if v.Next == v {
				// zhi you yige yuansu
			}
			v.Bkt = v.Bkt.Prev
			v.Prev = nil
			v.Next = v.Bkt.Children
			v.Bkt.Children = v
			v.Next.Prev = v
		} else {
			// only 1 child
			if v.Bkt.Children.Next == v.Bkt.Children {
				v.Bkt.Val++
			} else {
				// more than one element in this bucket
				// newBkt := &Bucket{
				// 	Val:      v.Bkt.Val + 1,
				// 	Prev:     v.Bkt.Prev,
				// 	Next:     v.Bkt,
				// 	Children: v,
				// }
				// if v.Bkt.Prev != nil {
				// 	v.Bkt.Prev.Next = newBkt
				// 	v.Bkt.Prev = newBkt
				// } else {
				// 	v.Bkt.Prev = newBkt
				// 	this.head = newBkt
				// }
				// if v.Bkt.Prev != nil {
				// 	v.Bkt.Prev.Next = v.Bkt.Next
				// }
				// if v.Bkt.Next != nil {
				// 	v.Bkt.Next.Prev = v.Bkt.Prev
				// }
				// v.Bkt = newBkt
			}

		}
	} else {
		n := &Node{
			Key: key,
		}
		this.hash[key] = n
		// add to tail
		if this.tail != nil && this.tail.Val == 1 {
			n.Bkt = this.tail
			n.Next = this.tail.Children
			n.Prev = this.tail.Children.Prev
			this.tail.Children.Prev.Next = n
			this.tail.Children.Prev = n
			this.tail.Children = n
		} else {
			// has no val == 1 bkt
			n.Prev = n
			n.Next = n
			newTailBkt := &Bucket{
				Val:      1,
				Children: n,
			}
			n.Bkt = newTailBkt
			if this.tail == nil {
				this.head = newTailBkt
				this.tail = newTailBkt
			} else {
				this.tail.Next = newTailBkt
				newTailBkt.Prev = this.tail
				this.tail = newTailBkt
			}
			if n.Key == "banana" {
				println("Here", this.head.Next.Children.Key)
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
