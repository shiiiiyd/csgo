package slicetest

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])

	s2 = append(s2, 2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

// 切片共享
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	t.Log(year, len(year), cap(year))

	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknown"
	t.Log(Q2)
	t.Log(year)
}

func TestSliceComparing(t *testing.T) {
	// 切片不能比较
	// a := []int{1, 2, 3, 4}
	// b := []int{1, 2, 3, 4}
	// if a == b {
	// 	t.Log("equal")
	// }

	// 数组可以比较
	var c [3]int
	d := [3]int{1, 2, 3}
	e := [...]int{1, 2, 3}
	t.Log(c)
	if d == e {
		t.Log("true")
	} else {
		t.Log("false")
	}
}
