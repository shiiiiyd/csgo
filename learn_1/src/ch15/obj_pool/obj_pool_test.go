package objpool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	t.Error(err)
	//}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Millisecond * 100); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")

}
