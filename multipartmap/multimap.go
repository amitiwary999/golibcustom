package multipartmap

import "hash/crc32"

type mapData[T comparable] struct {
	m map[T]interface{}
}

type partion[T comparable] struct {
	size int
	val  []*mapData[T]
}

func (p *partion[T]) findPartition(key string) int {
	hash := crc32.ChecksumIEEE([]byte(key))
	return int(hash) % p.size
}

func NewPartitionedMap[T comparable](size int) *partion[T] {
	partions := make([]*mapData[T], 0, size)
	for i := 0; i < size; i++ {
		m := make(map[T]interface{})
		partions = append(partions, &mapData[T]{m})
	}
	return &partion[T]{size: size, val: partions}
}
