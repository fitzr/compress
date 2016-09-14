package compress

import (
	"encoding/binary"
	"io"
)

func Pack(w io.Writer, name, ids []byte) (err error) {
	nameLen := uint32(len(name))
	idsLen := uint32(len(ids))

	err = binary.Write(w, binary.LittleEndian, nameLen)
	if err != nil {
		return
	}

	err = binary.Write(w, binary.LittleEndian, idsLen)
	if err != nil {
		return
	}

	_, err = w.Write(name)
	if err != nil {
		return
	}

	_, err = w.Write(ids)
	if err != nil {
		return
	}

	return
}

func Unpack(r io.Reader) (name, ids []byte, err error) {
	bytes, err := read(r, 4)
	if err != nil {
		return
	}
	nameLen := binary.LittleEndian.Uint32(bytes)

	bytes, err = read(r, 4)
	if err != nil {
		return
	}
	idsLen := binary.LittleEndian.Uint32(bytes)

	name, err = read(r, nameLen)
	if err != nil {
		return
	}

	ids, err = read(r, idsLen)
	if err != nil {
		return
	}

	return
}

func read(r io.Reader, size uint32) (bytes []byte, err error) {
	bytes = make([]byte, size)
	_, err = io.ReadFull(r, bytes)
	return bytes, err
}
