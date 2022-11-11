package arraytest

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr1, arr2)
	t.Log(arr[1], arr[2])
}
