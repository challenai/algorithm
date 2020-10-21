package solution

var resu []string
var ticket_ [][]string
var current []string
var path map[int]bool
var hash map[string][]int

func findItinerary(ticket [][]string) []string {
	resu = []string{}
	hash = map[string][]int{}
	ticket_ = ticket
	current = []string{"JFK"}
	path = map[int]bool{}
	for i := 0; i < len(ticket); i++ {
		if _, ok := hash[ticket[i][0]]; ok {
			hash[ticket[i][0]] = append(hash[ticket[i][0]], i)
		} else {
			hash[ticket[i][0]] = []int{i}
		}
	}
	dfs("JFK")
	return resu
}

func dfs(s string) {
	if len(current) == len(ticket_)+1 {
		if comparePath(current, resu) {
			resu = make([]string, len(current))
			copy(resu, current)
		}
		return
	}
	next, _ := hash[s]
	for i := 0; i < len(next); i++ {
		if _, ok := path[next[i]]; !ok {
			path[next[i]] = false
			current = append(current, ticket_[next[i]][1])
			dfs(ticket_[next[i]][1])
			current = current[:len(current)-1]
			delete(path, next[i])
		}
	}
}

func comparePath(p1, p2 []string) bool {
	if len(p2) == 0 {
		return true
	}
	for i := 0; i < len(p1); i++ {
		if p1[i] < p2[i] {
			return true
		}
		if p1[i] > p2[i] {
			return false
		}
	}
	return false
}
