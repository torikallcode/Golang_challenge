package codewars

import (
	"fmt"
	"strconv"
)

// func SumMixedArray(arr []any) int {
// 	sum := 0
// 	for _, i := range arr {
// 		switch v := i.(type) {
// 		case string:
// 			num, _ := strconv.Atoi(v)
// 			sum += num
// 		case int:
// 			sum += v
// 		}
// 	}
// 	return sum
// }

func SumMix(arr []any) int {
	sum := 0
	for _, i := range arr {
		v, _ := strconv.Atoi(fmt.Sprintf("%v", i))
		sum += v
	}
	return sum
}
