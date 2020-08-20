package bloom

import (
	"github.com/felipebool/bloom/pkg/filter"
	"github.com/felipebool/bloom/pkg/hash"
)

/**
@TODO calculate and return probability of a string be in the filter
https://en.wikipedia.org/wiki/Bloom_filter#Probability_of_false_positives

@TODO add tests
*/

type Bloom interface {
	Add(string)
	Check(string) bool
}

type bloom struct {
	filter filter.Filter
	Data   hash.Hash
}

func (b bloom) Add(s string) {
	b.Data.SetBits(b.filter.GetIndexes(s))
}

func (b bloom) Check(s string) bool {
	indexes := b.filter.GetIndexes(s)
	values := b.Data.GetValues(indexes)

	for _, value := range values {
		if !value {
			return false
		}
	}

	return true
}

func NewWithBoolSlice(size int, withCrypto bool) Bloom {
	return newBloom(hash.NewBoolHash(size), size, withCrypto)
}

func NewWithIntSlice(size int, withCrypto bool) Bloom {
	return newBloom(hash.NewIntHash(size), size, withCrypto)
}

func NewWithBigInt(size int, withCrypto bool) Bloom {
	return newBloom(hash.NewBigIntHash(), size, withCrypto)
}

func newBloom(d hash.Hash, size int, withCrypto bool) Bloom {
	b := bloom{Data: d}

	if withCrypto {
		b.filter = filter.NewCryptoFilter(size)
		return &b
	}

	b.filter = filter.NewNonCryptoFilter(size)
	return &b
}
