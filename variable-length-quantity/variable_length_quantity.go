package variablelengthquantity

import (
	"fmt"
	"math"
)

func DecodeVarint(bytes []byte) ([]uint32, error) {
	v := []byte{}
	result := []uint32{}
	for _, b := range bytes {
		v = append(v, b)
		if b < 0x80 {
			num := uint64(0)
			for _, bv := range v {
				num = num<<7 | uint64(bv&0x7f)
			}
			if uint64(num) <= math.MaxUint32 {
				result = append(result, uint32(num))
			} else {
				return nil, fmt.Errorf("overflow")
			}
			v = v[:0]
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("incomplete number")
	}
	return result, nil
}

func EncodeVarint(values []uint32) []byte {
	result := []byte{}
	for _, n := range values {
		v := []byte{byte(n & 0x7f)}
		n = n >> 7
		for n > 0 {
			v = append([]byte{byte(n&0x7f | 0x80)}, v...)
			n >>= 7
		}
		result = append(result, v...)
	}
	return result
}
