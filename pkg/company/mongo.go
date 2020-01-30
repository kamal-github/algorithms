package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
# A video game has recorded the player name and score for each time the game has been played.
# Create a high-score list of the top 50 Players and their best score, where each player only appears once on
# the high score list

# Input:
# Alice 50
# Joe 20
# Ben 65
# Alice 55
# Alice 65
# Joe 30
# Charlie 21
# ... (N entries)
#
# Output
# Ben 65
# Alice 65
# Joe 30
# Charlie 21
# ...(50 entries)
*/

// To execute Go code, please declare a func main() in a package "main"

type PlayerScore struct {
	name  string
	score int
}

// An IntHeap is a min-heap of ints.
type IntHeap []PlayerScore

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].score < h[j].score }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(PlayerScore))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	
	const MaxEntries = 6
	input := []string{"Alice 50", "Joe 20", "Ben 65", "Alice 55", "Alice 65", "Joe 30", "Charlie 21"}
	
	m := make(map[string]int)
	
	h := &IntHeap{}
	heap.Init(h)
	
	for _, str := range input {
		nameScores := strings.Split(str, " ")
		
		n := nameScores[0]
		sc, _ := strconv.Atoi(nameScores[1])
		
		if index, ok := m[n]; ok && (-sc) < (*h)[index].score {
			(*h)[index].score = -sc
		} else {
			heap.Push(h, PlayerScore{name: n, score: -sc})
			m[n] = h.Len() - 1
		}
		
	}
	
	for i := 0; i < MaxEntries; i++ {
		ps := heap.Pop(h).(PlayerScore)
		fmt.Println(ps.name, -ps.score)
	}
	
	sort.SliceStable()
}

