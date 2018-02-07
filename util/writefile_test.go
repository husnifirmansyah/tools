package util

import "testing"

func TestWriteFile(t *testing.T) {
	fa := &ForAll{
		Filename: "../files/firstlinear_sample.out",
	}
	numchar, err := fa.WriteFile("coba cek husni\n")
	if numchar != 15 {
		t.Errorf("got: %v, want: 15\n", numchar)
	}
	if err != nil {
		t.Error("error write file")
	}
}
