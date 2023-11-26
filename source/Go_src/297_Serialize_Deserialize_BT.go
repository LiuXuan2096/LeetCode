package Go_src

import (
	"strconv"
	"strings"
)

var SEP string = ","
var Null string = "#"

//type Codec struct {
//}
//
//func Constructor() Codec {
//
//	return Codec{}
//}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	sb := new(strings.Builder)
	serialize_helper(root, sb)
	return sb.String()
}

func serialize_helper(root *TreeNode, sb *strings.Builder) {
	if root == nil {
		sb.WriteString(Null)
		sb.WriteString(SEP)
		return
	}
	// 前序位置
	sb.WriteString(strconv.Itoa(root.Val))
	sb.WriteString(SEP)

	serialize_helper(root.Left, sb)
	serialize_helper(root.Right, sb)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if sp[0] == "#" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{val, build(), build()}
	}
	return build()
}
