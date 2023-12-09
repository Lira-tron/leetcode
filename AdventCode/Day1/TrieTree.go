package leetcode

type Node struct {
	Neighbors   map[byte]*Node
	Value       int
	IsEndOfWord bool
}

type TrieTree struct {
	r *Node
}

func (t *TrieTree) AddWord(w string, value int) {
	cur := t.r
	for i := 0; i < len(w); i++ {
		c := w[i]
		if n, ok := cur.Neighbors[c]; ok {
			cur = n
			continue
		}
		cur.Neighbors[c] = newNode()
		cur = cur.Neighbors[c]
	}
	cur.IsEndOfWord = true
	cur.Value = value
}

func (t *TrieTree) RemoveWord(w string) bool {
	found, _ := removeWordRec(w, 0, t.r)
	return found
}

func removeWordRec(w string, i int, n *Node) (bool, bool) {
	if i == len(w) {
		if !n.IsEndOfWord {
			return false, false
		}
		n.IsEndOfWord = false
		return true, len(n.Neighbors) == 0
	}
	c := w[i]
	next, ok := n.Neighbors[c]
	if !ok {
		return false, false
	}
	isFound, isDel := removeWordRec(w, i+1, next)
	if isDel {
		delete(n.Neighbors, c)
		return isFound, len(n.Neighbors) == 0 && !n.IsEndOfWord
	}
	return isFound, false
}

func (t *TrieTree) Find(w string) (int, bool) {
	cur := t.r
	for i := 0; i < len(w); i++ {
		c := w[i]
		if next, ok := cur.Neighbors[c]; ok {
			cur = next
		} else {
			return 0, false
		}
	}
	return cur.Value, cur.IsEndOfWord
}

func (t *TrieTree) FindFirst(w string) (int, bool) {
	cur := t.r
	for i := 0; i < len(w); i++ {
		if cur.IsEndOfWord {
			return cur.Value, cur.IsEndOfWord
		}
		c := w[i]
		if next, ok := cur.Neighbors[c]; ok {
			cur = next
		} else {
			return 0, false
		}
	}
	return cur.Value, cur.IsEndOfWord
}

func NewTrieTree() TrieTree {
	return TrieTree{
		r: newNode(),
	}
}

func newNode() *Node {
	return &Node{
		Neighbors: map[byte]*Node{},
	}
}
