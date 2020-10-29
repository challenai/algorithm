package solution

import (
	"testing"
)

func TestReadBinaryWatch(t *testing.T) {
	n := 1
	r := map[string]bool{"1:00": false, "2:00": false, "4:00": false, "8:00": false, "0:01": false, "0:02": false, "0:04": false, "0:08": false, "0:16": false, "0:32": false}
	rsu := readBinaryWatch(n)

	if len(r) != len(rsu) {
		t.Fail()
	}
	for i := 0; i < len(rsu); i++ {
		if _, ok := r[rsu[i]]; !ok {
			t.Fail()
		}
	}
}
