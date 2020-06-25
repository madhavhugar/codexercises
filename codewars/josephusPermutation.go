package main

import (
	"fmt"
	"sort"
)

func josephusPermutation(alive []interface{}, k int, dead []interface{}, begKill int) []interface{} {
	fmt.Println("before:alive", alive)
	fmt.Println("before:dead", dead)
	if len(alive) == 0 {
		return alive
	}
	indices := []int{}
	peopleBeforeRitual := len(alive)
	for i := (begKill % (len(alive))); i < peopleBeforeRitual; i = i + k {
		indices = append(indices, i)
	}
	fmt.Println("indices", indices)
	alive, justKilled := removeIndices(alive, indices)
	dead = append(dead, justKilled...)
	fmt.Println("after:dead", dead)
	fmt.Println("after:alive", alive)
	if len(alive) <= 1 {
		return append(dead, alive...)
	}
	// if len(alive) <= k {
	// 	nextKill := (k - ((peopleBeforeRitual - 1) - indices[0]) - 1) % (len(alive))
	// 	return josephusPermutation(alive, k, dead, nextKill)
	// }
	// nextBegKill := (indices[len(indices)-1] - (len(indices) - 1)) % (len(alive) - 1)
	fmt.Println("k", k, "(peopleBeforeRitual - 1)", (peopleBeforeRitual - 1), "indices[0]", indices[0])
	nextBegKill := k - ((peopleBeforeRitual - 1) - indices[0]) - 1

	fmt.Println("peopleBeforeRitual", peopleBeforeRitual)
	fmt.Println("nextBegKill", nextBegKill)
	return josephusPermutation(alive, k, dead, nextBegKill)
}

func removeByIndex(items []interface{}, index int) ([]interface{}, interface{}) {
	removed := items[index]
	items = append(items[:index], items[index+1:]...)
	return items, removed
}

func reverseSort(nums []int) {
	// in-place reverse sorting
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
}

func removeIndices(items []interface{}, indices []int) ([]interface{}, []interface{}) {
	reverseSort(indices)
	var final []interface{}
	var removedItems []interface{}
	for i := 0; i < len(indices); i++ {
		temp, removed := removeByIndex(items, indices[i])
		final = temp
		removedItems = append(removedItems, removed)
	}
	reversedRemoved := make([]interface{}, len(removedItems))
	for i, j := 0, len(removedItems)-1; i < len(removedItems); i, j = i+1, j-1 {
		reversedRemoved[i] = removedItems[j]
	}
	return final[:len(final)-len(reversedRemoved)+1], reversedRemoved
}
