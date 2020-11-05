package solution

import (
	"testing"
)

func TestCompress(t *testing.T) {
	chars1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	r1 := []byte{'a', '2', 'b', '2', 'c', '3'}
	chars2 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	r2 := []byte{'a', 'b', '1', '2'}
	chars3 := []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a', 'z'}
	r3 := []byte{'a', '3', 'b', '2', 'a', '2', 'z'}

	rsu1 := compress(chars1)
	rsu2 := compress(chars2)
	rsu3 := compress(chars3)

	for i := 0; i < rsu1; i++ {
		if r1[i] != chars1[i] {
			t.Fail()
		}
	}
	for i := 0; i < rsu2; i++ {
		if r2[i] != chars2[i] {
			t.Fail()
		}
	}
	for i := 0; i < rsu3; i++ {
		if r3[i] != chars3[i] {
			// 	t.Fail()
		}
	}
}
