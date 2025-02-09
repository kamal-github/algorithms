package main

import (
	"container/heap"
	"fmt"
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
	playerName string
	score      int // priority
}

type TopLeaderboard []PlayerScore

func (s TopLeaderboard) Len() int           { return len(s) }
func (s TopLeaderboard) Less(i, j int) bool { return s[i].score > s[j].score }
func (s TopLeaderboard) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *TopLeaderboard) Push(p any) {
	*s = append(*s, p.(PlayerScore))
}

func (s *TopLeaderboard) Pop() any {
	n := len(*s)
	last := (*s)[n-1]
	*s = (*s)[0 : n-1]
	return last
}

func main() {
	scores := []PlayerScore{
		{
			playerName: "A",
			score:      20,
		},
		{
			playerName: "B",
			score:      30,
		},
		{
			playerName: "D",
			score:      310,
		},
		{
			playerName: "C",
			score:      10,
		},
	}

	highestScores := make(map[string]int)
	for _, s := range scores {
		if oldScore, ok := highestScores[s.playerName]; ok && s.score > oldScore {
			highestScores[s.playerName] = s.score
			continue
		}
		highestScores[s.playerName] = s.score
	}

	topLeaderboard := make(TopLeaderboard, 0)
	heap.Init(&topLeaderboard)

	for pn, ps := range highestScores {
		heap.Push(&topLeaderboard, PlayerScore{
			playerName: pn,
			score:      ps,
		})
	}

	fmt.Println(topLeaderboard)

	fmt.Print(heap.Pop(&topLeaderboard))
	fmt.Print(heap.Pop(&topLeaderboard))
	fmt.Print(heap.Pop(&topLeaderboard))
	fmt.Print(heap.Pop(&topLeaderboard))
}
