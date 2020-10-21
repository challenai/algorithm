package recurse

import (
	"testing"
)

func TestArrage(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	arrResu := Arrange3(nums)
	for i := 0; i < len(arrResu); i++ {
		for j := 0; j < len(arrResu[i]); j++ {
			print(arrResu[i][j], " ")
		}
		println()
	}
	comResu := Combinate3(nums)
	for i := 0; i < len(comResu); i++ {
		for j := 0; j < len(comResu[i]); j++ {
			print(comResu[i][j], " ")
		}
		println()
	}
	t.Fail()
}
