package bloom_test

import (
	"github.com/felipebool/bloom/pkg/bloom"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBigIntBloomFilter(t *testing.T) {
	testCases := map[string]struct {
		add      string
		check    string
		expected bool
		cryptoHash bool
	}{
		"add equals to check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: true,
		},
		"add different from check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: true,
		},
		"add equals to check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: false,
		},
		"add different from check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithBigInt(100, tc.cryptoHash)
			bf.Add(tc.add)

			got := bf.Check(tc.check)

			require.Equal(t, tc.expected, got)
		})
	}
}

func TestIntBloomFilter(t *testing.T) {
	testCases := map[string]struct {
		add      string
		check    string
		expected bool
		cryptoHash bool
	}{
		"add equals to check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: true,
		},
		"add different from check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: true,
		},
		"add equals to check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: false,
		},
		"add different from check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithIntSlice(100, tc.cryptoHash)
			bf.Add(tc.add)

			got := bf.Check(tc.check)

			require.Equal(t, tc.expected, got)
		})
	}
}

func TestBoolBloomFilter(t *testing.T) {
	testCases := map[string]struct {
		add        string
		check      string
		expected   bool
		cryptoHash bool
	}{
		"add equals to check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: true,
		},
		"add different from check - crypto on": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: true,
		},
		"add equals to check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
			cryptoHash: false,
		},
		"add different from check - crypto off": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
			cryptoHash: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithBoolSlice(100, tc.cryptoHash)
			bf.Add(tc.add)

			got := bf.Check(tc.check)

			require.Equal(t, tc.expected, got)
		})
	}
}
