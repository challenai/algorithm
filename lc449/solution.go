package solution

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	str string
	tr  *TreeNode
}

func Constructor() Codec {
	return Codec{str: "", tr: nil}
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	var ptr *TreeNode
	q := []*TreeNode{root}
	elCnt := 1
	for elCnt > 0 {
		ptr = q[0]
		q = q[1:]
		if ptr == nil {
			q = append(q, nil)
			q = append(q, nil)
			this.str += "#,"
			continue
		}
		elCnt--
		this.str += strconv.Itoa(ptr.Val) + ","
		if ptr.Left != nil {
			elCnt++
		}
		if ptr.Right != nil {
			elCnt++
		}
		q = append(q, ptr.Left, ptr.Right)
	}
	return this.str
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return this.tr
	}
	nodes := []*TreeNode{}
	counter := -1
	// ptr := this.tr
	currentVal := ""
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			counter++
			if strings.Contains(currentVal, "#") {
				currentVal = ""
				nodes = append(nodes, nil)
				continue
			}
			v, err := strconv.Atoi(currentVal)
			if err != nil {
				panic("the input data is wrong")
			}
			nodes = append(nodes, &TreeNode{Val: v})
			if counter == 0 {
				// handle root, root has no parent
			} else if n := nodes[(counter-1)>>1]; n != nil {
				if counter%2 == 0 {
					n.Right = nodes[len(nodes)-1]
				} else {
					n.Left = nodes[len(nodes)-1]
				}
			} else {
				panic("the input data is wrong")
			}
			currentVal = ""
			continue
		}
		currentVal += string(data[i])
	}
	if len(nodes) > 0 {
		this.tr = nodes[0]
	}
	return this.tr
}
