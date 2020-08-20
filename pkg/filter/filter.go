package filter

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"hash/fnv"
	"log"
	"unsafe"
)

const intSize = int(unsafe.Sizeof(0))

type Filter interface {
	GetIndexes(string) []int
}

type filter struct {
	size int
	hashFunctions []func(str string) (uint32, error)
}

func (f filter) GetIndexes(str string) []int {
	result := make([]int, 4)

	for k, h := range f.hashFunctions {
		index, err := h(str)
		if err != nil {
			log.Fatalf("%+v", err)
		}

		result[k] = int(index) % f.size
	}

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

/**
	from: https://github.com/umahmood/GPSHash/blob/master/gpshash.go
 */
func hashPJW(str string) (uint32, error) {
	const maxUint = ^uint(0)
	const bitsPerWord = uint(32 * (1 + maxUint>>63))
	const threeQuarters = uint((bitsPerWord * 3) / 4)
	const oneEighth = uint(bitsPerWord / 8)
	const highBits = uint(1<<(bitsPerWord-oneEighth) - 1)

	var hash uint
	var test uint
	for i := 0; i < len(str); i++ {
		hash = (hash << oneEighth) + uint(str[i])
		test = hash & highBits
		if test != 0 {
			hash ^= (test >> threeQuarters) & (^highBits)
		}
	}

	return uint32(hash), nil
}

func hashMD5(str string) (uint32, error) {
	h := md5.New()
	byteString := []byte(str)

	return getResultAccordingToByteOrder(h.Sum(byteString)), nil
}

func hashSHA1(str string) (uint32, error) {
	h := sha1.New()
	byteString := []byte(str)

	return getResultAccordingToByteOrder(h.Sum(byteString)), nil
}

func hashSHA256(str string) (uint32, error) {
	h := sha256.New()
	byteString := []byte(str)

	return getResultAccordingToByteOrder(h.Sum(byteString)), nil
}

func getResultAccordingToByteOrder(value []byte) uint32 {
	if getEndian() {
		return binary.BigEndian.Uint32(value)
	}

	return binary.LittleEndian.Uint32(value)
}

/**
	hack from: https://github.com/virtao/GoEndian/blob/master/endian.go
	true, if it is big endian,
	false if it is little endian
*/
func getEndian() (ret bool) {
	var i = 0x1
	bs := (*[intSize]byte)(unsafe.Pointer(&i))

	return bs[0] == 0
}

func NewCryptoFilter(size int) Filter {
	return newFilter(size, getCryptoFunctions())
}

func NewNonCryptoFilter(size int) Filter {
	return newFilter(size, getNonCryptoFunctions())
}

func newFilter(size int, f []func(str string) (uint32, error)) Filter {
	return &filter{
		size: size,
		hashFunctions: f,
	}
}

func getCryptoFunctions() []func(str string) (uint32, error) {
	h := make([]func(str string) (uint32, error), 3)

	h[0] = hashSHA256
	h[1] = hashMD5
	h[2] = hashSHA1

	return h
}

func getNonCryptoFunctions() []func(str string) (uint32, error) {
	h := make([]func(str string) (uint32, error), 3)

	h[0] = hashFNV
	h[1] = hashFNVa
	h[2] = hashPJW

	return h
}
