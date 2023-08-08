package sort

func BubbleSort(array []int) {
	for i := range array {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] >= array[j+1] {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp
			}
		}
	}
}
