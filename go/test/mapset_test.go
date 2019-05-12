package test

import "testing"

func TestMApForSet(t *testing.T) {
	mySet:=map[int]bool{}
	mySet[0] = true
	n:=0
	if mySet[n] {
		t.Log(n," is existing")
	}else{
		t.Log(n," is not existing")
	}
	t.Log(len(mySet))
	delete(mySet,1)
	if mySet[n] {
		t.Log(n," is existing")
	}else{
		t.Log(n," is not existing")
	}
}
