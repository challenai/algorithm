package solution

var hash map[string][]int
var path map[int]bool
var current, rsu []string
var ticket_ [][]string

func findItinerary(ticket [][]string) []string {
	ticket_ = ticket
	hash = map[string][]int{}
	path = map[int]bool{}
	current = []string{"JFK"}
	rsu = []string{}
	for i := 0; i < len(ticket); i++ {
		if _, ok := hash[ticket[i][0]]; ok {
			hash[ticket[i][0]] = append(hash[ticket[i][0]], i)
		} else {
			hash[ticket[i][0]] = []int{i}
		}
	}
	dfs("JFK")
	// for i := 0; i < len(rsu); i++ {
	// 	print(rsu[i], " ")
	// }
	// println()
	return rsu
}

func dfs(s string) {
	if len(current) == len(ticket_)+1 {
		if comparePath(current, rsu) > 0 {
			rsu = make([]string, len(current))
			copy(rsu, current)
		}
		return
	}
	tks, _ := hash[s]
	for i := 0; i < len(tks); i++ {
		if _, ok := path[tks[i]]; !ok {
			path[tks[i]] = false
			current = append(current, ticket_[tks[i]][1])
			dfs(ticket_[tks[i]][1])
			delete(path, tks[i])
			current = current[:len(current)-1]
		}
	}
}

func comparePath(p1, p2 []string) int {
	if len(p2) == 0 {
		return 1
	}
	for i := 0; i < len(p1); i++ {
		if p1[i] < p2[i] {
			return 1
		}
		if p1[i] > p2[i] {
			return -1
		}
	}
	return 0
}
