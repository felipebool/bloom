package hash

type boolHash struct {
	data []bool
}

func (h boolHash) SetBits(indexes []int) {
	for _, index := range indexes {
		h.data[index] = true
	}
}

func (h boolHash) GetValues(indexes []int) []bool {
	result := make([]bool, len(indexes))

	for _, index := range indexes {
		result = append(result, h.data[index])
	}

	return result
}

func NewBoolHash(size int) Hash {
	return &boolHash{
		data: make([]bool, size),
	}
}
