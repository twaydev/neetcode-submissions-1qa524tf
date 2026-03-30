
type Node struct {
	children [26]*Node
	isEnd    bool
}

type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, char := range word {
		if node.children[char-'a'] == nil {
			node.children[char-'a'] = &Node{isEnd: false}
		}
		node = node.children[char-'a']
	}
	node.isEnd = true
}

func dfs(i int, node *Node, word string) bool {
	// kết thúc path
	if i == len(word) {
		return node.isEnd
	}
	// get char
	char := word[i]

	// special case with dots
	if char == '.' {
		for _, child := range node.children {
			if child != nil && dfs(i+1, child, word) {
				return true
			}
		}
		return false
	}

	// normal case without dots
	if node.children[char-'a'] == nil {
		return false
	}
	node = node.children[char-'a']
	return dfs(i+1, node, word)
}
func (this *WordDictionary) Search(word string) bool {
	return dfs(0, this.root, word)
}