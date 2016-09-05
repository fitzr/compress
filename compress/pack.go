package compress

import (
	"encoding/binary"
	"io"
)

func Pack(w io.Writer, word, tags []byte) {
	wordLen := uint32(len(word))
	tagsLen := uint32(len(tags))

	err := binary.Write(w, binary.LittleEndian, wordLen)
	check(err)

	err = binary.Write(w, binary.LittleEndian, tagsLen)
	check(err)

	_, err = w.Write(word)
	check(err)

	_, err = w.Write(tags)
	check(err)
}

func Unpack(r io.Reader) (word, tags []byte) {
	var wordLen uint32
	var tagsLen uint32
	err := binary.Read(r, binary.LittleEndian, &wordLen)
	check(err)

	err = binary.Read(r, binary.LittleEndian, &tagsLen)
	check(err)

	word = make([]byte, wordLen)
	tags = make([]byte, tagsLen)

	_, err = r.Read(word)
	check(err)

	_, err = r.Read(tags)
	check(err)

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
