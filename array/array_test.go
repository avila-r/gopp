package array_test

import (
	"testing"

	"github.com/avila-r/gopp/array"
)

func Test_Size(t *testing.T) {
	_ = array.Of("a", "b")


}

func Test_Add(t *testing.T) {
	slice := array.Of("a", "b", "c")

	expected := 3
	if size := len(slice); size != expected {
		t.Errorf("error: expected %v, got %v", expected, size)
	}

	array.Add(&slice, "b")

	expected++
	if size := len(slice); size != expected {
		t.Errorf("error: expected %v, got %v", expected, size)
	}
}
