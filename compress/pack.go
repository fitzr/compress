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
	var nameLen uint32
	var idsLen uint32
	err = binary.Read(r, binary.LittleEndian, &nameLen)
	if err != nil {
		return
	}

	err = binary.Read(r, binary.LittleEndian, &idsLen)
	if err != nil {
		return
	}

	name = make([]byte, nameLen)
	ids = make([]byte, idsLen)

	_, err = r.Read(name)
	if err != nil {
		return
	}

	_, err = r.Read(ids)
	if err != nil {
		return
	}

	return
}
