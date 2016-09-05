package compress

import (
	"encoding/binary"
	"io"
)

func Pack(w io.Writer, word, tags []byte) (err error) {
	wordLen := uint32(len(word))
	tagsLen := uint32(len(tags))

	err = binary.Write(w, binary.LittleEndian, wordLen)
	if err != nil {
		return
	}

	err = binary.Write(w, binary.LittleEndian, tagsLen)
	if err != nil {
		return
	}

	_, err = w.Write(word)
	if err != nil {
		return
	}

	_, err = w.Write(tags)
	if err != nil {
		return
	}

	return
}

func Unpack(r io.Reader) (word, tags []byte, err error) {
	var wordLen uint32
	var tagsLen uint32
	err = binary.Read(r, binary.LittleEndian, &wordLen)
	if err != nil {
		return
	}

	err = binary.Read(r, binary.LittleEndian, &tagsLen)
	if err != nil {
		return
	}

	word = make([]byte, wordLen)
	tags = make([]byte, tagsLen)

	_, err = r.Read(word)
	if err != nil {
		return
	}

	_, err = r.Read(tags)
	if err != nil {
		return
	}

	return
}
