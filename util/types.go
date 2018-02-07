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

type SecondLinear struct {
	Filename       string
	Testnum        int
	Key            []int
	FirstQueryNum  []int
	SecondQueryNum []int
	FirstQuery     map[int][]string
	SecondQuery    map[int][]string
}

type ForAll struct {
	Filename string
}
