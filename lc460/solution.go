package solution

type Node struct {
	K, V, F    int
	Prev, Next *Node
}

type LFUCache struct {
	freq, data              map[int]*Node
	minFreq, capacity, size int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		freq:     map[int]*Node{},
		data:     map[int]*Node{},
		capacity: capacity,
		size:     0,
	}
}

func (this *LFUCache) Get(key int) int {
	if n, ok := this.data[key]; ok {
		if n.Prev == n {
			if n.F == this.minFreq {
				this.minFreq++
			}
			delete(this.freq, n.F)
		}
		n.F++
		if f, ok2 := this.freq[n.F]; ok2 {
			n.Prev = f.Prev
			n.Next = f
			f.Prev.Next = n
			f.Prev = n
			this.freq[n.F] = n
		} else {
			this.freq[n.F] = n
			n.Prev = n
			n.Next = n
		}
		return n.V
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if n, ok := this.data[key]; ok {
		this.Get(key)
		n.V = value
		return
	}
	newNode := &Node{
		K: key,
		V: value,
		F: 1,
	}
	this.data[key] = newNode
	if this.size < this.capacity {
		this.size++
		if f, ok := this.freq[1]; ok {
			newNode.Prev = f.Prev
			newNode.Next = f
			f.Prev.Next = newNode
			f.Prev = newNode
			this.freq[1] = newNode
		} else {
			this.freq[1] = newNode
			newNode.Prev = newNode
			newNode.Next = newNode
		}
	} else {
		// need eviction
		// evict
		println(this.minFreq, newNode.K, newNode.V)
		println(this.freq[this.minFreq])
		// temp := this.freq[this.minFreq].Prev
		temp := this.freq[this.minFreq]
		delete(this.data, temp.K)
		if temp.Prev == temp {
			delete(this.freq, this.minFreq)
		} else {
			this.freq[this.minFreq] = temp.Next
			temp.Prev.Next = temp.Next
			temp.Next.Prev = temp.Prev
		}
		// add
		if n2, ok := this.freq[1]; ok {
			println(n2.V)
			println(newNode.V)
			newNode.Prev = n2.Prev
			newNode.Next = n2
			n2.Prev.Next = newNode
			n2.Prev = newNode
			this.freq[1] = newNode
		} else {
			this.freq[1] = newNode
			newNode.Prev = newNode
			newNode.Next = newNode
		}
		ptr := this.freq[1]
		for i := 0; i < 3; i++ {
			print(ptr.V, " ")
			ptr = ptr.Next
		}
	}
	this.minFreq = 1
}
