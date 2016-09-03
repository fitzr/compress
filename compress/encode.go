package compress

const (
	topBit     = 128
	byteMaxLen = 10
)

func Encode(n uint64) []byte {
	bytes := [byteMaxLen]byte{}
	i := byteMaxLen - 1
	for {
		bytes[i] = byte(n % topBit)
		if n < topBit {
			break
		}
		i--
		n /= topBit
	}
	bytes[byteMaxLen-1] += topBit
	return bytes[i:]
}

func Decode(bytes []byte) uint64 {
	n := uint64(0)
	for _, b := range bytes {
		v := uint64(b)
		if v < topBit {
			n = n*topBit + v
		} else {
			n = n*topBit + v - topBit
		}
	}
	return n
}
