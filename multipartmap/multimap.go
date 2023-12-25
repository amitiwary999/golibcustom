package multipartmap

import (
	"fmt"
	"hash/crc32"
)

type mapData[T comparable] struct {
	m map[T]interface{}
}

type partion[T comparable] struct {
	size int
	val  []*mapData[T]
}

func findPartition[T comparable](key T, size int) int {
	str := fmt.Sprintf("%v", key)
	hash := crc32.ChecksumIEEE([]byte(str))
	return int(hash) % size
}

func NewPartitionedMap[T comparable](size int) *partion[T] {
	partions := make([]*mapData[T], 0, size)
	/** traverse for 'size' number of times and create array with size 'size'. each item has a map with
	key type any comparable and value interface{}. */
	for i := 0; i < size; i++ {
		m := make(map[T]interface{})
		partions = append(partions, &mapData[T]{m})
	}
	return &partion[T]{size: size, val: partions}
}

/*
* get the possible partition number of a key, get the map from array using the partition number as index
then set the value in that map
*/
func (p *partion[T]) Set(key T, val interface{}) {
	partN := findPartition[T](key, p.size)
	mp := p.val[partN]
	mp.m[key] = val
}

/*
* get the possible partition number of a key, get the map from array using the partition number as index
then get the value of that key from map
*/
func (p *partion[T]) Get(key T) interface{} {
	partN := findPartition[T](key, p.size)
	fmt.Printf("part to read %d\n", partN)
	mp := p.val[partN]
	return mp.m[key]
}
