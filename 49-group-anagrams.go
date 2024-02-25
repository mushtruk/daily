package main

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

func GroupAnagrams(strs []string) [][]string {
	strsToAnagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := SortString(str)
		if anagrams, found := strsToAnagrams[sortedStr]; found {
			strsToAnagrams[sortedStr] = append(anagrams, str)
		} else {
			strsToAnagrams[sortedStr] = []string{str}
		}
	}

	var result [][]string
	for _, anagrams := range strsToAnagrams {
		result = append(result, anagrams)
	}

	return result
}
