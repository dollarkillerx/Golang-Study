package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	updateSlice(s)
	fmt.Println(s)

	s1 := s[3:5]
	fmt.Println(s1)
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d",
		s, len(s), cap(s))

	arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 = arr[2:6]
	s2 := s1[3:5]
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 11)
	fmt.Println()
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))
	fmt.Println(s5, len(s5), cap(s5))
	// s4 and s5 no longer view arr.
	fmt.Println(arr)

}

/**
slice 可向后扩展 是数组的视图
len() slice总长度
cap() 不可超越数组长度
添加元素如果超越cap,系统会重新分配更大的底层数组
*/
