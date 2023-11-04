package twoFA

type Key struct {
	raw    []byte
	digits int
	offset int // offset of counter
}
