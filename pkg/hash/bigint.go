package hash

import "math/big"

type bigIntHash struct {
	data big.Int
}

func (h bigIntHash) SetBits(indexes []int) {
	for _, index := range indexes {
		h.data.SetBit(&h.data, index, 1)
	}
}

func (h bigIntHash) GetValues(indexes []int) []bool {
	result := make([]bool, len(indexes))

	for _, index := range indexes {
		result = append(result, h.data.Bit(index) == 1)
	}

	return result
}

func NewBigIntHash() Hash {
	return &bigIntHash{
		data: big.Int{},
	}
}
