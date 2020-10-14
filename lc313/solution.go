package solution

import "sort"

var product, maxPrime int
var primes_, potientialFactorSet []int

func nthSuperUglyNumber(n int, primes []int) int {
	resu := 1
	sort.Slice(primes, func(i, j int) bool {
		return primes[i] < primes[j]
	})
	if n == 1 {
		return resu
	}
	n--
	maxPrime = primes[len(primes)-1]
	potientialFactorSet = []int{}
	// potientialFactorSet = append(potientialFactorSet, primes...)
	product = 1
	primes_ = primes
	dfs(0)
	sort.Slice(potientialFactorSet, func(i, j int) bool {
		return potientialFactorSet[i] < potientialFactorSet[j]
	})
	// for i := 0; i < len(potientialFactorSet); i++ {
	// 	print(potientialFactorSet[i], " ")
	// }
	ptr := 0
	base := 1
	for n > 0 {
		resu = potientialFactorSet[ptr] * base
		println(resu)
		ptr++
		if ptr == len(potientialFactorSet) {
			ptr = 0
			base = maxPrime
		}
		n--
	}
	return resu
}

func dfs(startIdx int) {
	if product < maxPrime {
		potientialFactorSet = append(potientialFactorSet, product)
	} else {
		return
	}
	for i := startIdx; i < len(primes_); i++ {
		product *= primes_[i]
		dfs(i)
		product /= primes_[i]
	}
}
