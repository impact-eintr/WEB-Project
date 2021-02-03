package process_test

import (
	"Json/process"
	"testing"
)

func sub1(t *testing.T) {
	x := 1
	y := 2
	res1 := 3

	res2 := process.Add(x, y)
	if res2 != res1 {
		t.Errorf("%d\n", res2)
	}
}
func sub2(t *testing.T) {
	x := 1
	y := 2
	res1 := 3

	res2 := process.Add(x, y)
	if res2 != res1 {
		t.Errorf("%d\n", res2)
	}
}
func sub3(t *testing.T) {
	x := 1
	y := 2
	res1 := 3

	res2 := process.Add(x, y)
	if res2 != res1 {
		t.Errorf("%d\n", res2)
	}
}

func TestSub(t *testing.T) {
	t.Run("A=1", sub1)
	t.Run("A=2", sub2)
	t.Run("B=1", sub3)
}
