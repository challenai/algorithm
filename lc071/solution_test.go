package solution

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	inputs := []string{
		"/home/",
		"/../",
		"/home//foo/",
		"/a/./b/../../c/",
		"/a/../../b/../c//.//",
		"/a//b////c/d//././/..",
	}
	resu := []string{
		"/home",
		"/",
		"/home/foo",
		"/c",
		"/c",
		"/a/b/c",
	}

	r := []string{}
	var str string
	for i := 0; i < len(inputs); i++ {
		str = simplifyPath(inputs[i])
		r = append(r, str)
	}

	pass := true
	for i := 0; i < len(resu); i++ {
		if resu[i] != r[i] {
			pass = false
			break
		}
	}

	if !pass {
		t.Fail()
	}
}
