package filter

import (
	"crypto/md5"
	"crypto/sha1"
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
}

func (f filter) GetIndexes(str string) []int {
	result := make([]int, 4)

	index, err := hashFNV(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result[0] = int(index) % f.size

	index, err = hashFNVa(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result[1] = int(index) % f.size

	index, err = hashMD5(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result[2] = int(index) % f.size

	index, err = hashSHA1(str)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	// @FIXME this can be a problem
	result[3] = int(index) % f.size

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

func hashMD5(str string) (uint32, error) {
	h := md5.New()
	byteString := []byte(str)

	if getEndian() {
		return binary.BigEndian.Uint32(h.Sum(byteString)), nil
	}

	return binary.LittleEndian.Uint32(h.Sum(byteString)), nil
}

func hashSHA1(str string) (uint32, error) {
	h := sha1.New()
	byteString := []byte(str)

	if getEndian() {
		return binary.BigEndian.Uint32(h.Sum(byteString)), nil
	}

	return binary.LittleEndian.Uint32(h.Sum(byteString)), nil
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

func NewFilter(size int) Filter {
	return &filter{
		size: size,
	}
}

