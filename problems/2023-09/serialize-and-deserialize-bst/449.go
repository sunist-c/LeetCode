package serialize_and_deserialize_bst

import "encoding/json"

type OutputStruct struct {
	Val  int    `json:"val"`
	Path string `json:"path"`
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var output []OutputStruct
	if root != nil {
		output = step(root, []OutputStruct{}, "")
	}

	if data, err := json.Marshal(output); err == nil {
		return string(data)
	}

	return ""
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	var output []OutputStruct
	if err := json.Unmarshal([]byte(data), &output); err != nil {
		return nil
	}

	if len(output) == 0 {
		return nil
	}

	var root = &TreeNode{}
	for _, v := range output {
		plug(root, v)
	}

	return root
}

func step(t *TreeNode, input []OutputStruct, path string) (output []OutputStruct) {
	output = append(input, OutputStruct{Val: t.Val, Path: path})
	if t.Left != nil {
		output = step(t.Left, output, path+"l")
	}
	if t.Right != nil {
		output = step(t.Right, output, path+"r")
	}

	return output
}

func plug(root *TreeNode, node OutputStruct) {
	if node.Path == "" {
		root.Val = node.Val
		return
	}

	var current = root
	for _, v := range node.Path {
		if v == 'l' {
			if current.Left == nil {
				current.Left = &TreeNode{}
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = &TreeNode{}
			}
			current = current.Right
		}
	}
	current.Val = node.Val
}
