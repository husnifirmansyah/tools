package util

import "testing"

func TestFirstLinearWriteFile(t *testing.T) {
	fl := &FirstLinear{
		Filename: "../files/firstlinear_sample.out",
	}
	numchar, err := fl.WriteFile("coba cek husni\n")
	if numchar != 15 {
		t.Errorf("got: %v, want: 15\n", numchar)
	}
	if err != nil {
		t.Error("error write file")
	}
}
