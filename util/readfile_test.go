package util

import (
	"reflect"
	"testing"
)

func TestFirstLinearReadFile(t *testing.T) {
	fl := &FirstLinear{
		Filename: "../files/firstlinear_sample.in",
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

func TestSecondLinearReadFile(t *testing.T) {
	sl := &SecondLinear{
		Filename: "../files/secondlinear_sample.in",
	}

	sl.ReadFile()
	if sl.Testnum != 2 {
		t.Errorf("got: %v, want: 2\n", sl.Testnum)
	}
}

func TestThirdLinearReadFile(t *testing.T) {
	tl := &ThirdLinear{
		Filename: "../files/thirdlinear_sample.in",
	}
	tl.ReadFile()
	if tl.Testnum != 5 {
		t.Errorf("got: %v, want: 5\n", tl.Testnum)
	}
	query := []string{"0.25 1.0 0.1 0.01 0.5", "0.25 1.0 0.1 0.01 0.9", "0.00001 10000 0.00001 0.00001 1000", "0.4 10000 0.00001 0.00001 700", "1 100 1 1 10"}
	if !reflect.DeepEqual(tl.Query, query) {
		t.Errorf("got query: %v, want query: %v\n", tl.Query, query)
	}
}

func TestForthLinearReadFile(t *testing.T) {
	fo := &ForthLinear{
		Filename: "../files/forthlinear_sample.in",
	}
	fo.ReadFile()
	if fo.Testnum != 2 {
		t.Errorf("got: %v, want: 2\n", fo.Testnum)
	}

	vectornum := []int{3, 5}
	if !reflect.DeepEqual(fo.VectorNum, vectornum) {
		t.Errorf("got vector num: %v, want vector num: %v\n", fo.VectorNum, vectornum)
	}

	firstVector := []string{"1 3 -5", "1 2 3 4 5"}
	if !reflect.DeepEqual(fo.FirstVector, firstVector) {
		t.Errorf("got first vector: %v, want first vector: %v\n", fo.FirstVector, firstVector)
	}

	secondVector := []string{"-2 4 1", "1 0 1 0 1"}
	if !reflect.DeepEqual(fo.SecondVector, secondVector) {
		t.Errorf("got second vector: %v, want second vector: %v\n", fo.SecondVector, secondVector)
	}
}
