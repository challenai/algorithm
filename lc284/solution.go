package solution

type Iterator struct {
	data []int
	idx  int
}

func MockIterator() *Iterator {
	return &Iterator{
		data: []int{1, 2, 3},
		idx:  0,
	}
}

func (it *Iterator) Next() int {
	if it.idx >= len(it.data) || it.idx < 0 {
		return 0
	}
	it.idx++
	return it.data[it.idx-1]
}

func (it *Iterator) HasNext() bool {
	return it.idx < len(it.data)
}

type PeekingIterator struct {
	it    *Iterator
	data  int
	isEnd bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	temp := &PeekingIterator{
		it:    iter,
		isEnd: !iter.HasNext(),
	}
	if !temp.isEnd {
		temp.data = iter.Next()
	}
	return temp
}

func (this *PeekingIterator) HasNext() bool {
	return !this.isEnd
}

func (this *PeekingIterator) Next() int {
	if this.isEnd {
		return 0
	}
	this.isEnd = !this.it.HasNext()
	temp := this.data
	if !this.isEnd {
		this.data = this.it.Next()
	}
	return temp
}

func (this *PeekingIterator) Peek() int {
	return this.data
}
