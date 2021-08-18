package answer

func CheckNumberFrequency(inputs []int) ([]int, map[int]int) {
	mapResults := getFrequency(inputs)
	keys, values := arrayToMap(mapResults)

	// order key and value, based on values
	sortBasedOnValues(0, len(values), keys, values)

	// order key
	orderKey(keys, values)
	return keys, mapResults
}

func getFrequency(inputs []int) map[int]int {
	result := make(map[int]int)
	for _, v := range inputs {
		result[v]++
	}
	return result
}

func arrayToMap(inputs map[int]int) (keys, values []int) {
	keys = make([]int, 0)
	values = make([]int, 0)
	for k, v := range inputs {
		keys = append(keys, k)
		values = append(values, v)
	}
	return
}

func sortBasedOnValues(start, end int, keys []int, values []int) {
	finish := false
	for !finish {
		finish = true
		for i := start; i < end-1; i++ {
			if values[i] < values[i+1] {
				Swap(&keys[i], &keys[i+1])
				Swap(&values[i], &values[i+1])
				finish = false
			}
		}
	}
}

func orderKey(keys []int, values []int) {
	freqs := getFrequency(values)
	for k, v := range freqs {
		if v > 1 {
			index := findIndex(values, k)
			customSort(index, index+v, keys)
		}
	}
}

func findIndex(inputs []int, value int) int {
	for k, v := range inputs {
		if v == value {
			return k
		}
	}

	return 0
}

func customSort(start, end int, inputs []int) {
	finish := false
	for !finish {
		finish = true
		for i := start; i < end-1; i++ {
			if inputs[i] > inputs[i+1] {
				Swap(&inputs[i], &inputs[i+1])
				finish = false
			}
		}
	}
}
