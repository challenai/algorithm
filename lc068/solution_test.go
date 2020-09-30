package solution

import (
	"testing"
)

func TestFullJustify(t *testing.T) {

	words1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	maxWidth1 := 16
	r1 := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}
	words2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	maxWidth2 := 16
	r2 := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}
	words3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}
	maxWidth3 := 20
	r3 := []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}
	resu1 := fullJustify(words1, maxWidth1)
	resu2 := fullJustify(words2, maxWidth2)
	resu3 := fullJustify(words3, maxWidth3)

	if len(resu1) != len(r1) || len(resu2) != len(r2) || len(resu3) != len(r3) {
		t.Fail()
	}
}
