package codewars

func ReverseSeq(n int) []int {

	var arr []int
	for i := n; i > 0; i-- {
		arr = append(arr, i)
	}

	return arr
}
