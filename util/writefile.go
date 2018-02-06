package util

import (
	"bufio"
	"os"
)

func (fl *FirstLinear) WriteFile(contents string) (numchar int, err error) {
	f, _ := os.Create(fl.Filename)
	w := bufio.NewWriter(f)
	numchar, err = w.WriteString(contents)
	w.Flush()
	return
}
