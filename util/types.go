package util

type FileUtil interface {
	ReadFile()
}

type FirstLinear struct {
	Filename  string
	Testnum   int
	Enginenum []int
	Querynum  []int
	Engine    map[int][]string
	Query     map[int][]string
}
