package test

import "testing"

func TestCompareArray(t *testing.T) {
	a:=[...]int{1,2,3,4}
	//b:=[...]int{1,2,3,4,5}
	//c:=[...]int{1,2,3,4,5,6}
	d:=[...]int{1,2,3,4}
	//t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestMap(t *testing.T) {
	m1 := make(map[int]string,10)
	t.Log(len(m1))
}