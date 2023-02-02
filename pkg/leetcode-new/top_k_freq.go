package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]NumToFreq)
	for _, n := range nums {
		if v, ok := freqMap[n]; ok {
			incFreq := v.freq + 1
			freqMap[n] = NumToFreq{num: n, freq: incFreq}
		} else {
			freqMap[n] = NumToFreq{num: n, freq: 1}
		}
	}

	var numToFreqList NumToFreqList
	for _, v := range freqMap {
		numToFreqList = append(numToFreqList, v)
	}

	sort.Slice(numToFreqList, func(i, j int) bool {
		return numToFreqList[i].freq > numToFreqList[j].freq
	})

	topKElements := make([]int, k)

	for i := 0; i < k; i++ {
		topKElements[i] = numToFreqList[i].num
	}

	return topKElements
}

type NumToFreq struct {
	num, freq int
}

type NumToFreqList []NumToFreq

//func (nf NumToFreqList) Less(i, j int) bool {
//	// decreasing order
//	return nf[i].freq > nf[j].freq
//}
//
//func (nf NumToFreqList) Len() int {
//	return len(nf)
//}
//
//func (nf NumToFreqList) Swap(i, j int) {
//	nf[i], nf[j] = nf[j], nf[i]
//}
