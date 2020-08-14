package hash

type intHash struct {
	data []int
}

func (h intHash) SetBits(indexes []int) {
	for _, index := range indexes {
		h.data[index] = 1
	}
}

func (h intHash) GetValues(indexes []int) []bool {
	result := make([]bool, len(indexes))

	for _, index := range indexes {
		result = append(result, h.data[index] == 1)
	}

	return result
}

func NewIntHash(size int) Hash {
	return &intHash{
		data: make([]int, size),
	}
}
