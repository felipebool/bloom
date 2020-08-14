package filter

import (
	"hash/fnv"
	"log"
)

type Filter interface {
	GetIndexes(string) []int
}

type filter struct {
	size int
}

func (f filter) GetIndexes(str string) []int {
	result := make([]int, 3)

	index, err := hashFNV(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result = append(result, int(index) % f.size)

	index, err = hashFNVa(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result = append(result, int(index) % f.size)

	return result
}

func hashFNV(str string) (uint32, error) {
	n := fnv.New32()

	if _, err := n.Write([]byte(str)); err != nil {
		return 0, err
	}

	return n.Sum32(), nil
}

func hashFNVa(str string) (uint32, error) {
	n := fnv.New32a()

	if _, err := n.Write([]byte(str)); err != nil {
		return 0, err
	}

	return n.Sum32(), nil
}

func NewFilter(size int) Filter {
	return &filter{
		size: size,
	}
}

