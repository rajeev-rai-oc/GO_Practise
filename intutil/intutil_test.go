package main

import (
	"testing"
)

// A test is created by writing a function with a name
// beginning with `Test`.
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// `t.Error*` will report test failures but continue
		// executing the test. `t.Fatal*` will report test
		// failures and stop the test immediately.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}
