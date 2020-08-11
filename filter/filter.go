package filter

import "math/big"

/**
	@TODO calculate and return probability of a string be in the filter
	https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives

	@TODO add tests
 */

type Bloom interface {
	AddString(str string) error
	CheckString(str string) (bool, error)
}

type bloom struct {
	size      uint32
	functions []func(str string) (uint32, error)
	hashMap   big.Int
}

func (b *bloom) AddString(str string) error {
	for _, f := range b.functions {
		index, err := f(str)
		if err != nil {
			return err
		}

		b.hashMap.SetBit(&b.hashMap, int(index), 1)
	}

	return nil
}

func (b *bloom) CheckString(str string) (bool, error) {
	for _, f := range b.functions {
		index, err := f(str)
		if err != nil {
			return false, err
		}

		if b.hashMap.Bit(int(index)) == 0 {
			return false, nil
		}
	}

	return true, nil
}

func NewBloom(size uint32, functions []func(str string) (uint32, error)) *bloom {
	return &bloom{
		size: size,
		functions: functions,
		hashMap: big.Int{},
	}
}
