package main

// https://leetcode.com/problems/merge-strings-alternately/description/?envType=study-plan-v2&envId=leetcode-75

func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)

	nw := ""
	i := 0
	for ; i < l1 && i < l2; i++ {
		nw += string(word1[i]) + string(word2[i])
	}

	if i < l1 {
		nw += word1[i:]
	}
	if i < l2 {
		nw += word2[i:]
	}
	return nw
}
