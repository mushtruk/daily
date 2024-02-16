package main

// boyer-moore voting algorithm
// it works unless there are majority element appears at least len/2
// otherwise second

func MajorityElement(nums []int) int {
	var candidate, count int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if candidate == num {
			count++
		} else {
			count--
		}
	}

	// optional unles count < len/2
	// count = 0
	// for _, num := range nums {
	// 	if candidate == num {
	// 		count++
	// 	}
	// }

	// if count > len(nums)/2 {
	// 	return candidate
	// }
	// return -1

	return candidate
}
