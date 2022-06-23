package math

import (
	"fmt"
	"math"
)

func GetPrime() {
	print("请输入一个大于2的整数: ")
	var max int
	_, err := fmt.Scan(&max)
	if err != nil || max < 2 {
		return
	}

	for i := 2; i <= max; i++ {
		if isPrime(i) {
			fmt.Printf("%d 是质数\n", i)
		}
	}
}

func isPrime(n int) bool {
	edge := int(math.Sqrt(float64(n)))
	for i := 2; i <= edge; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
