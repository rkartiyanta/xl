package answer

import "fmt"

func Concat(a, b string) string {
	lengthA := len(a)
	lengthB := len(b)

	max := lengthA
	if lengthB < lengthA {
		max = lengthB
	}
	var result string
	for i := 0; i < max; i++ {
		result = fmt.Sprintf("%s%c%c", result, a[i], b[i])
	}

	if lengthA > max {
		result = fmt.Sprintf("%s%s", result, a[max:lengthA])
	}

	if lengthB > max {
		result = fmt.Sprintf("%s%s", result, b[max:lengthB])
	}
	return result
}
