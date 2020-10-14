package solution

import (
	"strconv"
)

func getHint(secret, guess string) string {
	if len(secret) != len(guess) {
		return "0A0B"
	}
	hash := map[byte]int{}
	for i := 0; i < len(secret); i++ {
		if _, ok := hash[secret[i]]; ok {
			hash[secret[i]]++
		} else {
			hash[secret[i]] = 1
		}
	}
	numA := 0
	numB := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			numA++
			hash[secret[i]]--
		}
	}
	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			continue
		}
		if v, ok := hash[guess[i]]; ok && v > 0 {
			hash[guess[i]]--
			numB++
		}
	}
	return strconv.Itoa(numA) + "A" + strconv.Itoa(numB) + "B"
}
