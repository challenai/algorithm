package solution

import "strconv"

var path map[string]bool
var matrix_ [][]byte
var word_ string
var resu bool

func exist(matrix [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	matrix_ = matrix
	word_ = word
	resu = false
	path = map[string]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == word[0] {
				key := serialize(i, j)
				path[key] = false
				searchPath(i, j, 1)
				delete(path, key)
			}
		}
	}
	return resu
}

func searchPath(i, j, startIdx int) {
	if resu {
		return
	}
	if startIdx == len(word_) {
		resu = true
		return
	}
	var key string
	if i > 0 && matrix_[i-1][j] == word_[startIdx] {
		key = serialize(i-1, j)
		if _, ok := path[key]; !ok {
			path[key] = false
			searchPath(i-1, j, startIdx+1)
			delete(path, key)
		}
	}
	if i < len(matrix_)-1 && matrix_[i+1][j] == word_[startIdx] {
		key = serialize(i+1, j)
		if _, ok := path[key]; !ok {
			path[key] = false
			searchPath(i+1, j, startIdx+1)
			delete(path, key)
		}
	}
	if j > 0 && matrix_[i][j-1] == word_[startIdx] {
		key = serialize(i, j-1)
		if _, ok := path[key]; !ok {
			path[key] = false
			searchPath(i, j-1, startIdx+1)
			delete(path, key)
		}
	}
	if j < len(matrix_[0])-1 && matrix_[i][j+1] == word_[startIdx] {
		if _, ok := path[key]; !ok {
			key = serialize(i, j+1)
			path[key] = false
			searchPath(i, j+1, startIdx+1)
			delete(path, key)
		}
	}
}

func serialize(i, j int) string {
	return strconv.Itoa(i) + "_" + strconv.Itoa(j)
}
