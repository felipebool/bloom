package function

import "hash/fnv"

func HashFNV() func(str string) (uint32, error) {
	return func(str string) (uint32, error) {
		n := fnv.New32()

		if _, err := n.Write([]byte(str)); err != nil {
			return 0, err
		}

		return n.Sum32(), nil
	}
}

func HashFNVa() func(str string) (uint32, error) {
	return func(str string) (uint32, error) {
		n := fnv.New32a()

		if _, err := n.Write([]byte(str)); err != nil {
			return 0, err
		}

		return n.Sum32(), nil
	}
}
