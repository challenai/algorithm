package solution

// use a bucket replace plain Node to get a O(1) element move
// when exists many elements with same value
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
	if v, ok := this.hash[key]; ok {
		// remove v to prev bucket
		// find original bucket
	} else {
		this.hash[key] = &Node{
			Key: key,
		}
		// this.hash[key] = &Node{
		// 	Key: key,
		// }
		// this.head = this.hash[key]
		// this.head.Next.Prev = this.head
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
