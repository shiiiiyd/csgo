package auto_growing

import "testing"

const NumOfElems = 10000
const Times = 1000

func TestAutoGrow(t *testing.T) {
	for i := 0; i < Times; i++ {
		var s []int
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func TestProperInit(t *testing.T) {
	for i := 0; i < Times; i++ {
		s := make([]int, 0, NumOfElems)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAutoGrow(b *testing.B) {
	b.ResetTimer()
	var s []int
	for i := 0; i < b.N; i++ {
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
	b.StopTimer()
}

func BenchmarkProperInit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
	b.StopTimer()
}

func BenchmarkOverSizeInit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, NumOfElems*8)
		for j := 0; j < NumOfElems; j++ {
			s = append(s, j)
		}
	}
	b.StopTimer()
}
