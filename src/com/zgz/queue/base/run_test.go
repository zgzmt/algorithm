package base

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	got := Run([]int{1, 2, 3, 4, 5, 6})
	want := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("excepted:%v,got:%v", want, got)
	}
}
