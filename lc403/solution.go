package solution

var rsu bool
var units map[int]bool

func canCross(nums []int) bool {
	// idea: use dfs to search all possible path,
	// cut branch if it unable to go forward
	if len(nums) < 2 {
		return true
	}
	if nums[1] != 1 || nums[0] != 0 {
		return false
	}
	units = map[int]bool{}
	for i := 1; i < len(nums)-1; i++ {
		units[nums[i]] = false
	}
	units[nums[len(nums)-1]] = true
	rsu = false
	dfs(1, 1)
	println(rsu)
	return rsu
}

func dfs(unit, step int) {
	println(unit, step)
	if rsu {
		return
	}
	if r, ok := units[unit]; ok {
		if r {
			rsu = true
			return
		}
		if step > 1 {
			dfs(unit+step-1, step-1)
		}
		dfs(unit+step, step)
		dfs(unit+step+1, step+1)
	}
}
