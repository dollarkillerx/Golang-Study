package _func

import (
	"math/rand"
	"testing"
)

func returnMultiValues() (int,int){
	return rand.Intn(10),rand.Intn(20)
}

func TestFn(t *testing.T) {
	a,b := returnMultiValues()
	t.Log(a,b)
}
