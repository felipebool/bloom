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
	}{
		"add equals to check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
		},
		"add different from check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithBigInt(100)
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
	}{
		"add equals to check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
		},
		"add different from check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithIntSlice(100)
			bf.Add(tc.add)

			got := bf.Check(tc.check)

			require.Equal(t, tc.expected, got)
		})
	}
}

func TestBoolBloomFilter(t *testing.T) {
	testCases := map[string]struct {
		add      string
		check    string
		expected bool
	}{
		"add equals to check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "It’s been a long day, and the last thing you want to do is start",
			expected: true,
		},
		"add different from check": {
			add:      "It’s been a long day, and the last thing you want to do is start",
			check:    "\"It’s been a long day, and the last",
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			bf := bloom.NewWithIntSlice(100)
			bf.Add(tc.add)

			got := bf.Check(tc.check)

			require.Equal(t, tc.expected, got)
		})
	}
}