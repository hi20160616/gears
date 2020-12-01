package gears

import (
	"reflect"
	"testing"
)

func checkStrSlices(t *testing.T, _got, _want []string) {
	t.Helper()
	if !reflect.DeepEqual(_got, _want) {
		t.Errorf("\n[Got]\n%v\n[Want]\n%v\n", _got, _want)
	}
}

func TestExists(t *testing.T) {
	checkExists := func(t *testing.T, _got, _want bool) {
		t.Helper()
		if _got != _want {
			t.Errorf("\nGot\n%v\nWant\n%v\n", _got, _want)
		}
	}
	t.Run("test if true while file exist.", func(t *testing.T) {
		got := Exists("./io.go")
		want := true
		checkExists(t, got, want)
	})
	t.Run("test if false while file not exist.", func(t *testing.T) {
		got := Exists("./io_false.go")
		want := false
		checkExists(t, got, want)
	})
	t.Run("test if true while path exist.", func(t *testing.T) {
		got := Exists("./")
		want := true
		checkExists(t, got, want)
	})
	t.Run("test if true while path not exist.", func(t *testing.T) {
		got := Exists("./testwrongpath/")
		want := false
		checkExists(t, got, want)
	})
}

func TestGetUnPrefixedFiles(t *testing.T) {
	t.Run("test get filelist has prefix [04.23]", func(t *testing.T) {
		got, _ := GetPrefixedFiles("./test/", "[04.25]")
		want := []string{"test/[04.25][0323H]test.txt"}
		checkStrSlices(t, got, want)
	})
}
func TestGetPrefixedFiles(t *testing.T) {
	t.Run("test get filelist has prefix [04.25]", func(t *testing.T) {
		got, _ := GetPrefixedFiles("./test/", "[04.25]")
		want := []string{"test/[04.25][0323H]test.txt"}
		checkStrSlices(t, got, want)
	})
	t.Run("test get filelist has prefix [04.23]", func(t *testing.T) {
		got, _ := GetPrefixedFiles("./test/", "[04.23]")
		want := []string{"test/[04.23][1323H]test.txt", "test/[04.23][2335H]test.txt"}
		checkStrSlices(t, got, want)
	})
}
