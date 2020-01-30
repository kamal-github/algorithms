package leetcode

// https://leetcode.com/problems/prefix-and-suffix-search/

type TrieNode struct {
	m         map[int32]*TrieNode
	weights   []int
	endOfPath bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		m:         make(map[int32]*TrieNode),
		weights:   make([]int, 0),
		endOfPath: false,
	}
}

func (tn *TrieNode) Insert(s string, i int) {
	trav := tn
	for _, ch := range s {
		if next, ok := trav.m[ch]; ok {
			trav.weights = append(trav.weights, i)
			trav = next
		} else {
			t := NewTrieNode()
			trav.m[ch] = t
			trav.weights = append(trav.weights, i)
			trav = t
		}
	}
	
	trav.endOfPath = true
}

// return stored weight, if prefix matched
func (tn *TrieNode) Search(prefix string) []int {
	trav := tn
	var prev TrieNode
	for _, ch := range prefix {
		if next, ok := trav.m[ch]; ok {
			trav = next
			prev = *trav
		} else {
			return []int{}
		}
	}
	
	
		return prev.weights
	//return []int{}
}

type WordFilter struct {
	words   []string
	preTrie *TrieNode
	sufTrie *TrieNode
}

func Constructor(words []string) WordFilter {
	preTrie := NewTrieNode()
	for i, w := range words {
		preTrie.Insert(w, i)
	}
	
	suffTrie := NewTrieNode()
	for i, w := range words {
		suffTrie.Insert(Reverse(w), i)
	}
	
	return WordFilter{words: words, preTrie:preTrie, sufTrie:suffTrie}
}

func (wf *WordFilter) F(prefix string, suffix string) int {
	
	return -1
}

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}
