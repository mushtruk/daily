package main

// var (
// 	ransomNote = "aa"
// 	magazine   = "aab"
// )

func CanConstruct(ransomNote string, magazine string) bool {
	magazineCount := make(map[rune]int)

	for _, char := range magazine {
		magazineCount[char]++
	}

	for _, char := range ransomNote {
		if mc, ok := magazineCount[char]; !ok || mc == 0 {
			return false
		}
		magazineCount[char]--
	}
	return true
}
