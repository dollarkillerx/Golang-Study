package unit

import "testing"

func TestGetSum(t *testing.T) {
	sum := GetSum(10)
	if sum != 100 {
		t.Errorf("期望%d,实际%d\n",100,sum)
		t.FailNow()
	}
	t.Log("TestGetSum 测试成功")
}

func TestGetSum2(t *testing.T) {
	sum := GetSum(10)
	if sum != 100 {
		t.Fatal("我是一个致命错误")
	}
	t.Log("TestGetSum 测试成功")
}