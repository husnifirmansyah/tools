package util

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	fl := &FirstLinear{
		Filename: "sample.in",
	}
	fl.ReadFile()
	if fl.Testnum != 2 {
		t.Errorf("got: %v, want: 2\n", fl.Testnum)
	}
	if !reflect.DeepEqual(fl.Enginenum, []int{5, 5}) {
		t.Errorf("got: %v, want: %v\n", fl.Enginenum, []int{5, 5})
	}
	if !reflect.DeepEqual(fl.Querynum, []int{10, 7}) {
		t.Errorf("got: %v, want: %v\n", fl.Querynum, []int{10, 7})
	}
	engine := map[int][]string{
		0: []string{"Yeehaw", "NSM", "Dont Ask", "B9", "Googol"},
		1: []string{"Yeehaw", "NSM", "Dont Ask", "B9", "Googol"},
	}
	if !reflect.DeepEqual(fl.Engine, engine) {
		t.Errorf("got: %v, want: %v\n", fl.Engine, engine)
	}
	query := map[int][]string{
		0: []string{"Yeehaw", "Yeehaw", "Googol", "B9", "Googol", "NSM", "B9", "NSM", "Dont Ask", "Googol"},
		1: []string{"Googol", "Dont Ask", "NSM", "NSM", "Yeehaw", "Yeehaw", "Googol"},
	}
	if !reflect.DeepEqual(fl.Query, query) {
		t.Errorf("got: %v, want: %v\n", fl.Query, query)
	}
}
