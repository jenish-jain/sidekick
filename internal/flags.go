package internal

type BooleanFlag struct {
	name         string
	shortHand    string
	defaultValue bool
	usage        string
}

func NewBooleanFlag(name string, shortHand string, defaultVal bool, usage string) *BooleanFlag {
	return &BooleanFlag{
		name:         name,
		shortHand:    shortHand,
		defaultValue: defaultVal,
		usage:        usage,
	}
}

func (f *BooleanFlag) Name() string {
	return f.name
}

func (f *BooleanFlag) ShortHand() string {
	return f.shortHand
}

func (f *BooleanFlag) DefaultValue() bool {
	return f.defaultValue
}

func (f *BooleanFlag) Usage() string {
	return f.usage
}
