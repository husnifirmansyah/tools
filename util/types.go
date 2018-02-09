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

type ThirdLinear struct {
	Filename string
	Testnum  int
	Query    []string
}

type ForthLinear struct {
	Filename     string
	Testnum      int
	VectorNum    []int
	FirstVector  []string
	SecondVector []string
}

type ForAll struct {
	Filename string
}
