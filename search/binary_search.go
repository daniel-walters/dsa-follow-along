package search

func BinarySearch(haystack []int, needle int) bool {
	lo := 0             // inclusive
	hi := len(haystack) //exclusive

	for lo < hi {
		mid := lo + (hi-lo)/2
		v := haystack[mid]

		if v == needle {
			return true
		} else if needle > v {
			lo = mid + 1 // inclusive
		} else {
			hi = mid // exclusive
		}
	}

	return false
}
