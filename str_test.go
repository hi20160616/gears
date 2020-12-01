package gears

import (
	"fmt"
	"testing"

	"github.com/axgle/mahonia"
)

func TestStrSliceDiff(t *testing.T) {
	s1 := []string{"1", "5", "2", "3"}
	s2 := []string{"4", "2", "3"}

	want := []string{"1", "5"}
	got2 := StrSliceDiff2(s1, s2)
	got := StrSliceDiff(s1, s2)
	if got[0] != want[0] || got[1] != want[1] || got2[0] != want[0] || got2[1] != want[1] {
		t.Errorf("\nGot: %v, Want %v\n", got, want)
	}
}

func TestConvertToUtf8(t *testing.T) {
	enc := mahonia.NewEncoder("gbk")
	ts := enc.ConvertString("美国总统特朗普宣布将从德国撤离部分美国士兵")
	want := "美国总统特朗普宣布将从德国撤离部分美国士兵"
	err := ConvertToUtf8(&ts, "gbk", "utf-8")
	if err != nil {
		t.Errorf("ConvertToUtf8() invoked error: %v", err)
	}
	if want != ts {
		t.Errorf("want: %v, got: %v", want, ts)
	}
	fmt.Println(ts)
}
