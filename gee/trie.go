package gee

import "strings"

//动态路由通常具备功能
//参数匹配
//通配

type node struct {
	pattern string  //待匹配路由例如/p/：lang
	part    string  //路由中的一部分例如:lang
	chidren []*node //子节点
	iswild  bool    //是否精确匹配 ,part含有:或者*时为true
}

// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.chidren {
		if child.part == part || child.iswild {
			return child
		}
	}
	return nil
}

// 素有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.chidren {
		if child.part == part || child.iswild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, iswild: part[0] == ':' || part[0] == '*'}
		n.chidren = append(n.chidren, child)
	}
	child.insert(pattern, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	chirden := n.matchChildren(part)

	for _, child := range chirden {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
