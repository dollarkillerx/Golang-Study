package unit

import "errors"

func GetSum(n int) int {
	var sum = 0
	if n<0 {
		panic(errors.New("n必须大于0"))
	}
	for i := 1;i < n+1;i++ {
		sum += n
	}
	return sum
}
