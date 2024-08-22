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

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
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
	params := strings.Split(pathname, "/")

	result := dfs(tn, params)

	if result == nil {
		return nil
	}
	return result.Handler
}

func dfs(node *TreeNode, params []string) *TreeNode {
	currentParam := params[0]
	isLastParam := len(params) == 1

	for _, child := range node.Children {
		if isLastParam {
			if isGeneral(child.Param) {
				return child
			}

			if child.Param == currentParam {
				return child
			}

			continue
		}

		if !isGeneral(child.Param) && child.Param != currentParam {
			continue
		}

		result := dfs(child, params[1:])

		if result != nil {
			return result
		}
	}

	return nil
}
