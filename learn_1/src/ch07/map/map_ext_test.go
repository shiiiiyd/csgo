package map_ext_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(opt int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 3)

	if mySet[1] {
		t.Log("1 is existing")
	} else {
		t.Log("1 is not existing")
	}

	t.Log(len(mySet))
}
