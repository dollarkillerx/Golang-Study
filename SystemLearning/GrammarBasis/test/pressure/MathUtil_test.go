package unit

import "testing"

func BenchmarkGetSum(b *testing.B) {
	b.Log("Benchmark Start")
	//汇报内存开销
	b.ReportAllocs()
	for i:=0;i<b.N ;i++  {
		GetSum(100)
	}
}
