package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 110
}

func mm() {
	s := [10]int{0,1,2,3,4,5,6,7,8,9}
	s1 := s[2:6] // 2,3,4,5
	s2 := s1[4:6] // 6,7
	fmt.Println(s2)
}

//func slice() {
//	var s []int // zero value for slice is nil
//	//s == nil
//	for i:=0;i<50;i++ {
//		s = append()
//	}
//}

func main()  {
	arr := [...]int{0,1,2,3,4,5,6,7,8,9}
	a := arr[2:6]
	b := arr[2:] // 从2去到最后
	c := arr[:6] //  从头开始取6个
	d := arr[:] // 全部
	updateSlice(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	mm()
	fmt.Printf("arr= %v len(arr)= %d cap(arr)=%d",a,len(a),cap(a));
	fmt.Println(' ')
	s1 := append(a,10)
	fmt.Println(s1)
	fmt.Println(arr)

	s2 := make([]int,16,20)
	fmt.Println("value" , s2 , len(s2), cap(s2))

	copy(s2,s1)
	fmt.Println(s2)

}