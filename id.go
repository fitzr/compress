package compress

import (
	"bytes"
	"strconv"
)

func CompressIds(idStrings []string) []byte {
	idNumbers := toIntArray(idStrings)
	toDiff(idNumbers)
	return encode(idNumbers)
}

func DecompressIds(idBytes []byte) []string {
	idNumbers := decode(idBytes)
	restoreFromDiff(idNumbers)
	return toStringArray(idNumbers)
}

func toDiff(idNumbers []uint64) {
	for i := len(idNumbers) - 1; i > 0; i-- {
		idNumbers[i] -= idNumbers[i-1]
	}
}

func restoreFromDiff(idNumbers []uint64) {
	for i := 1; i < len(idNumbers); i++ {
		idNumbers[i] += idNumbers[i-1]
	}
}

func toIntArray(strings []string) []uint64 {
	numbers := make([]uint64, len(strings))
	for i, v := range strings {
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			panic(err)
		}
		numbers[i] = n
	}
	return numbers
}

func toStringArray(numbers []uint64) []string {
	strings := make([]string, len(numbers))
	for i, v := range numbers {
		strings[i] = strconv.FormatUint(v, 10)
	}
	return strings
}

func encode(numbers []uint64) []byte {
	buff := new(bytes.Buffer)
	for _, id := range numbers {
		Encode(buff, id)
	}
	return buff.Bytes()
}

func decode(data []byte) []uint64 {
	numbers := make([]uint64, 0)
	r := bytes.NewReader(data)
	for {
		n, err := Decode(r)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				panic(err)
			}
		}
		numbers = append(numbers, n)
	}
	return numbers
}
