package client

import (
	"github.com/shiiiiyd/csgo/go_learning/src/ch11/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(10))
}
