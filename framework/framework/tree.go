package framework

import (
	"net/http"
	"strings"
)

type TreeNode struct {
	Children []*TreeNode
	Handler  func(rw http.ResponseWriter, r *http.Request)
	Param    string
}

func NewTreeNode() TreeNode {
	return TreeNode{
		Children: []*TreeNode{},
		Param:    "",
	}
}

func (tn *TreeNode) Insert(pathname string, handler func(rw http.ResponseWriter, r *http.Request)) {
	node := tn

	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.FindChild(param)
		if child == nil {
			child = &TreeNode{
				Children: []*TreeNode{},
				Param:    param,
			}
			node.Children = append(node.Children, child)
		}
		node = child
	}
	node.Handler = handler
}

func (tn *TreeNode) FindChild(param string) *TreeNode {
	for _, child := range tn.Children {
		if child.Param == param {
			return child
		}
	}
	return nil
}

func (tn *TreeNode) Search(pathname string) func(rw http.ResponseWriter, r *http.Request) {
	node := tn

	params := strings.Split(pathname, "/")

	for _, param := range params {
		child := node.FindChild(param)
		if child == nil {
			return nil
		}
		node = child
	}
	return node.Handler
}
