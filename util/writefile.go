package util

import (
	"bufio"
	"os"
)

func (fa *ForAll) WriteFile(contents string) (numchar int, err error) {
	f, _ := os.Create(fa.Filename)
	w := bufio.NewWriter(f)
	numchar, err = w.WriteString(contents)
	w.Flush()
	return
}
