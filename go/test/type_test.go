package test

import "testing"

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Logf("%T,%T",a,aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*"+s+"*")
}