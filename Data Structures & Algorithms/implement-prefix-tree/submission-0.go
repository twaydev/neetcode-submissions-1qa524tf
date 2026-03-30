
type Node struct {
	children [26]*Node
	isEnd    bool
}

type PrefixTree struct {
	root *Node
}

func Constructor() PrefixTree {
	return PrefixTree{root: &Node{}}
}

func (this *PrefixTree) Insert(word string) {
	node := this.root
	for _, ch := range word {
		if node.children[ch-'a'] == nil {
			node.children[ch-'a'] = &Node{isEnd: false}
		}
		node = node.children[ch-'a']
	}
	node.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		if node.children[ch-'a'] == nil {
			return false
		}
		node = node.children[ch-'a']
	}
	return node.isEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		if node.children[ch-'a'] == nil {
			return false
		}
		node = node.children[ch-'a']
	}
	return true
}