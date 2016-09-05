package compress

import (
	"bytes"
	"strconv"
)

func CompressIds(idStrings []string) []byte {
	// from string to unsigned int 64
	idNumbers := make([]uint64, len(idStrings))
	for i, v := range idStrings {
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			panic(err)
		}
		idNumbers[i] = n
	}

	// to first number and diff numbers
	for i := len(idNumbers) - 1; i > 0; i-- {
		idNumbers[i] -= idNumbers[i-1]
	}

	// encode
	buff := new(bytes.Buffer)
	for _, id := range idNumbers {
		Encode(buff, id)
	}
	return buff.Bytes()
}

func DecompressIds(idBytes []byte) []string {
	idNumbers := make([]uint64, 0)
	r := bytes.NewReader(idBytes)

	// decode
	for {
		n, err := Decode(r)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				panic(err)
			}
		}
		idNumbers = append(idNumbers, n)
	}

	// from diff to actual number
	for i := 1; i < len(idNumbers); i++ {
		idNumbers[i] += idNumbers[i-1]
	}

	// from unsigned int 64 to string
	idStrings := make([]string, len(idNumbers))
	for i, v := range idNumbers {
		idStrings[i] = strconv.FormatUint(v, 10)
	}
	return idStrings
}
