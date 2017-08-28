package test

import (
	"testing"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(10)
}

func TestTest(t *testing.T)  {
	t.Parallel()
	test()
}

func TestTest2(t *testing.T) {
	t.Parallel()
	test2()
}

