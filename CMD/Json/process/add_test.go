package process_test

import (
	"Json/process"
	"testing"
)

func TestAdd(t *testing.T) {
	x := 1
	y := 2
	res1 := 3

	res2 := process.Add(x, y)
	if res2 != res1 {
		t.Errorf("%d\n", res2)
	}
	t.Log(res2)
}
