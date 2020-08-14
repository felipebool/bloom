package hash

type Hash interface {
	SetBits(indexes []int)
	GetValues(indexes []int) []bool
}
