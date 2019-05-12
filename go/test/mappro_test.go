package test

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[0] = func(op int)int{return op}
	m[1] = func(op int)int {return op*op}
	m[2] = func(op int) int {return op*op*op}

	t.Log(m[0](2))
	t.Log(m[1](2))
	t.Log(m[2](2))
}