package main

// https://leetcode.com/explore/challenge/card/april-leetcoding-challenge-2021/593/week-1-april-1st-april-7th/3699/

var vowelsMap map[string]bool

func init() {
	vowelsMap = make(map[string]bool)
	vowelsMap = map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
}

func alikeStringHalves(s string) bool {
	mid := len(s) / 2
	a := s[0:mid]
	b := s[mid:]

	aNum := vowelCountMap(a)
	bNum := vowelCountMap(b)

	return aNum == bNum
}

func vowelCountMap(s string) int {
	count := 0
	for _, c := range s {
		if _, ok := vowelsMap[string(c)]; ok {
			count++
		}
	}
	return count
}
