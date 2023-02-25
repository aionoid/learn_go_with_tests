package arraysandslices

// @Title
// @Description
// @Author
// @Update

func Sum(array []int) (sum int) {
	// for i := 0; i < len(array); i++ {
	//   sum += array[i]
	// }
	// sum := 0
	for _, number := range array {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	out := make([]int, len(numbersToSum))
	for i, val := range numbersToSum {
		out[i] = Sum(val)
	}
	return out
}

func SumAll2(numbersToSum ...[]int) []int {
	var out []int
	for _, val := range numbersToSum {
		out = append(out, Sum(val))
	}
	return out
}

func SumAllTails(numbersToSum ...[]int) []int {
	out := make([]int, len(numbersToSum))
	for i, val := range numbersToSum {
		if len(val) == 0 {
			out[i] = 0
		} else {
			out[i] = Sum(val[1:])
		}
	}
	return out
}
