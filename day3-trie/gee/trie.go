package gee

import "strings"

type node struct {
	pattern  string  // 完整的路由，只有叶子节点才有值
	part     string  // 当前这一段的路由
	children []*node // 孩子节点
	isWild   bool    // 当前这一段是否模糊匹配了
}

// 匹配孩子，找到想要的那个，如果不存在就返回nil
func (n *node) matchChild(part string) *node {
	// 遍历孩子节点，找到第一个匹配的节点
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			return ch
		}
	}
	return nil
}

// 匹配孩子，找到所有匹配的节点，返回列表
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, ch := range n.children {
		if ch.part == part || ch.isWild {
			nodes = append(nodes, ch)
		}
	}
	return nodes
}

func (n *node) insert(height int, pattern string, patterns []string) {
	// 插入某个节点，从头节点开始，需要高度, 需要插入的节点的路由, 为了方便起见，还有路由的拆分成的组数
	if height == len(patterns) {
		// 插入结束
		n.pattern = pattern // 只有叶子节点才有值
		return
	}

	part := patterns[height] // 当前高度是0，对应的就是数组下标为0的元素，在当前的节点的子孩子们选择合适的地点插入
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		n.children = append(n.children, child)
	}
	child.insert(height+1, pattern, patterns)
}

func (n *node) search(height int, patterns []string) *node {
	if height == len(patterns) || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := patterns[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(height+1, patterns)
		if result != nil {
			return result
		}
	}
	return nil
}
