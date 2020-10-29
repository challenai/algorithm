package solution

import "strconv"

func fizzBuzz(n int) []string {
	resu := []string{}
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			resu = append(resu, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			resu = append(resu, "Fizz")
			continue
		}
		if i%5 == 0 {
			resu = append(resu, "Buzz")
			continue
		}
		resu = append(resu, strconv.Itoa(i))
	}
	return resu
}
