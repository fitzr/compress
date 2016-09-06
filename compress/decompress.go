package compress

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func Decompress(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 1024 * 256)
	writer := bufio.NewWriterSize(w, 1024 * 256)
	for decompressLine(reader, writer) {
	}
	writer.Flush()
}

func decompressLine(r io.Reader, w *bufio.Writer) bool {
	name, idBytes, err := Unpack(r)
	if err != nil {
		log.Println("unpack line error : ", err)
		return false
	}
	tags := DecompressIds(idBytes)
	w.Write(name)
	w.WriteByte('\t')
	w.WriteString(strings.Join(tags, ","))
	w.WriteByte('\n')
	return true
}
