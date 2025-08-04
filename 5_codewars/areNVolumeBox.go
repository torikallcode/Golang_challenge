package codewars

func GetSize(w, h, d int) [2]int {
	// your code here
	volume := 2 * ((w * h) + (w * d) + (h * d))
	luas := w * h * d
	arr := [2]int{volume, luas}
	return [2]int(arr)
}

// It("Test 4,2,6",func() {Expect(GetSize(4,2,6)).To(Equal([2]int{88,48}))})
//     It("Test 1,1,1",func() {Expect(GetSize(1,1,1)).To(Equal([2]int{6,1}))})
//     It("Test 1,2,1",func() {Expect(GetSize(1,2,1)).To(Equal([2]int{10,2}))})
//     It("Test 1,2,2",func() {Expect(GetSize(1,2,2)).To(Equal([2]int{16,4}))})
//     It("Test 10,10,10",func() {Expect(GetSize(10,10,10)).To(Equal([2]int{600,1000}))})
