package gopp_test

import (
	"testing"

	. "github.com/avila-r/gopp"
)

func Test_Condition(t *testing.T) {
	target := "a"
	if v := If(true, target).Else("b"); v != target {
		t.Errorf("error - expected %v, got %v", target, v)
	}

	target = "b"
	if v := If(false, "a").Else(target); v != target {
		t.Errorf("error - expected %v, got %v", target, v)
	}
}
