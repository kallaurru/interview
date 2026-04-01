package algos

type FixtureConfig struct {
	Data     []int
	Expected []int
	Idx      int
	K        int
	Ok       bool
}

type StringFixture struct {
	Input  string
	Output string
	Ok     bool
}
