package hashgo

import (
	"math/bits"
)

/** MurMurHash use multiply and xor to create hash from array byte */

func MurMurHash3(key []byte, seed uint32) uint32 {
	var length int = len(key)
	var blocks int = length / 4 // divide in 32 bits or 4 bytes blocks
	const mulc1 = 0xcc9e2d51
	const mulc2 = 0x1b873593
	ans := seed

	for i := 0; i < blocks; i++ {
		pbyt := uint32(key[i]) + uint32(key[i*4+1])<<8 + uint32(key[i*4+2])<<16 + uint32(key[i*4+3])<<24
		pbyt *= mulc1
		pbyt = bits.RotateLeft32(pbyt, 15)
		pbyt *= mulc2
		ans ^= pbyt
		pbyt = bits.RotateLeft32(pbyt, 13)
		ans = ans*5 + 0xe6546b64
	}
	tail := key[4*blocks:]
	var k1 uint32
	/** if length is multiple of 4 then it means when we work on block of 4 there will be some element left.
	  depending on how many element left we handle it*/
	switch length & 3 {
	case 3:
		k1 ^= uint32(tail[2]) << 16
		fallthrough

	case 2:
		k1 ^= uint32(tail[1]) << 8
		fallthrough
	case 1:
		k1 ^= uint32(tail[0])
		k1 *= mulc1
		k1 = bits.RotateLeft32(k1, 15)
		k1 *= mulc2
		ans ^= k1
	}

	ans ^= ans >> 16
	ans *= 0x85ebca6b
	ans ^= ans >> 13
	ans *= 0xc2b2ae35
	ans ^= ans >> 16
	return ans
}
