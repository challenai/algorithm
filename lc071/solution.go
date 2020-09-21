package solution

import "strings"

func simplifyPath(path string) string {
	resuList := []string{}
	for i := 0; i < len(path); i++ {
		if path[i] == '.' {
			if i > 0 && path[i-1] == '.' {
				if len(resuList) != 0 {
					resuList = resuList[:len(resuList)-1]
				}
				if len(resuList) != 0 {
					resuList = resuList[:len(resuList)-1]
				}
			} else if i > 0 && path[i-1] != '/' && path[i-1] != '.' {
				resuList = append(resuList, string(path[i]))
			}
		} else if path[i] == '/' {
			if len(resuList) == 0 || resuList[len(resuList)-1] != "/" {
				resuList = append(resuList, string(path[i]))
			}
		} else {
			resuList[len(resuList)-1] += string(path[i])
		}
	}
	if len(resuList) > 0 && resuList[len(resuList)-1] == "/" {
		resuList = resuList[:len(resuList)-1]
	}
	if len(resuList) == 0 {
		return "/"
	}
	return strings.Join(resuList, "")
}
