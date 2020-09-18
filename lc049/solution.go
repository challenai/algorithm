package solution

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	resultsHash := map[string][]string{}
	resu := [][]string{}
	for i := 0; i < len(strs); i++ {
		if _, ok := resultsHash[sortStr(strs[i])]; ok {
			resultsHash[sortStr(strs[i])] = append(resultsHash[sortStr(strs[i])], strs[i])
		} else {
			resultsHash[sortStr(strs[i])] = []string{strs[i]}
		}
	}

	for _, strsSet := range resultsHash {
		resu = append(resu, strsSet)
	}
	return resu
}

func sortStr(str string) string {
	temp := strings.Split(str, "")
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	return strings.Join(temp, "")
}
