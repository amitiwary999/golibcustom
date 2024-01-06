package hashgo

import "math/bits"

func MurMurHash(key []byte, seed uint32) uint32 {
	var length int = len(key)
	var blocks int = length / 4 // divide in 32 bits or 4 bytes blocks
	const mulc1 = 0xca8e2d41
	const mulc2 = 0xa9a92e51
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
