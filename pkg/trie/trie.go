package trie

type Trie struct {
	children    map[byte]*Trie
	endOfBranch bool
}

func (t *Trie) Insert(word string) {
	if t == nil {
		return
	}

	trav := t

	for i := 0; i < len(word); i++ {
		n, ok := trav.children[word[i]]
		if ok {
			trav = n
			continue
		}

		newNode := &Trie{
			children:    make(map[byte]*Trie),
			endOfBranch: false,
		}
		//Add node and continue
		if trav != nil {
			trav.children[word[i]] = newNode
		}
		// update trav
		trav = newNode
	}

	if trav != nil {
		trav.endOfBranch = true
	}
}

func (t *Trie) StartsWith(prefix string) bool {
	if t == nil {
		return false
	}

	trav := t

	for i := 0; i < len(prefix); i++ {
		n, ok := trav.children[prefix[i]]
		if ok {
			trav = n
		} else {
			return false
		}
	}

	return true
}

func (t *Trie) GetStringsStartsWith(prefix string) []string {
	if t == nil {
		return []string{}
	}

	trav := t

	// Reaching to a node where prefix completely matched
	for i := 0; i < len(prefix); i++ {
		n, ok := trav.children[prefix[i]]
		if ok {
			trav = n
		} else {
			return []string{}
		}
	}

	// if you are here, means you have found the prefix, and trav is at prefix end
	// get all the strings in a map of trie node pointer by trav
	words := make([]string, 0)
	getAllStrings(trav, prefix, &words)

	return words

}

func getAllStrings(t *Trie, tempWord string, words *[]string) {
	if t.endOfBranch == true {
		*words = append(*words, tempWord)
	}

	for b, tc := range t.children {
		getAllStrings(tc, tempWord+string(b), words)
	}
}

func (t *Trie) Search(word string) bool {
	if t == nil {
		return false
	}

	trav := t

	for i := 0; i < len(word); i++ {
		n, ok := trav.children[word[i]]
		if ok {
			trav = n
		} else {
			return false
		}
	}

	return trav.endOfBranch
}

//func (t *Trie) Delete(word string) error {
//
//}
