package algos

type FixtureConfig struct {
	Data     []int
	Expected []int
	Idx      int
	K        int
	Ok       bool
}

type StringFixtureAdv struct {
	L      string
	R      string
	Expect string
	Ok     bool
}

type StringFixture struct {
	Input  string
	Output string
	Ok     bool
}

type PaarFixture struct {
	Ok       bool
	Left     []int
	Right    []int
	Expected []int // size m+n
	M        int
	N        int
}

type TreeFixture struct {
	Ok         bool
	Raw        []int
	L          int
	H          int
	Expect     int
	ExpectBool bool
	MV         int
}

type CacheFixtureInt struct {
	Ok     bool
	Cmd    string
	Key    int
	Val    int
	Expect int
	SetId  int
}
