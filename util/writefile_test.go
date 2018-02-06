package util

import "testing"

func TestWriteFile(t *testing.T) {
	fl := &FirstLinear{
		Filename: "sample.out",
	}
	numchar, err := fl.WriteFile("coba cek husni\n")
	if numchar != 15 {
		t.Errorf("got: %v, want: 15\n", numchar)
	}
	if err != nil {
		t.Error("error write file")
	}
}
