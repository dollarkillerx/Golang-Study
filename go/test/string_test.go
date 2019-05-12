package test

import (
	"strconv"
	"strings"
	"testing"
)

func TestSString(t *testing.T) {
	s := "HELLO"
	t.Log(s[0])
}

func TestStringFn(t *testing.T) {
	s:="A,b,c"
	paras := strings.Split(s,",")
	t.Log(paras)

	t.Log(strings.Join(paras,"-"))
}

func TestConv(t *testing.T) {
	s:=strconv.Itoa(10)
	t.Logf("%T",s)
	if i,err :=strconv.Atoi("10");err==nil{
		t.Logf("%T",i)
	}
}
