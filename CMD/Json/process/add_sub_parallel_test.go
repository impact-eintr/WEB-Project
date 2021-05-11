package process_test

import (
	"testing"
	"time"
)

func parallelTest1(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
func parallelTest2(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func parallelTest3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

func TestSubParallel(t *testing.T) {
	t.Logf("SetUp")

	t.Run("group", func(t *testing.T) {
		t.Run("A=1", parallelTest1)
		t.Run("A=2", parallelTest2)
		t.Run("B=1", parallelTest3)
	})
	t.Logf("teardown")
}
