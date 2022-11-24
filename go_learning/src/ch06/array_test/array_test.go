package arraytest

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr)
	t.Log(len(arr), cap(arr))
	arr1 := []int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}

	t.Log(len(arr1), cap(arr1))
	t.Log(len(arr2), cap(arr2))

	t.Log(arr1, arr2)
	t.Log(arr[1], arr[2])

	arr3 := arr2[:]
	t.Log(arr3)
}
