package map_ext_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{}
	t.Log(m)
	t.Log(m["one"])

	m1 := map[string]int{"one": 1, "two": 2, "three": 3, "second": 4}
	t.Log(m1["one"])
	t.Logf("len m1=%d", len(m1))

	m2 := map[string]int{}
	m2["three"] = 17
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[string]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	t.Log(m1[3])
	if v, ok := m1[3]; ok {
		t.Logf("key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m1 := map[int]int{1: 1, 2: 4, 3: 9}

	for k, v := range m {
		t.Log(k, v)
	}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
