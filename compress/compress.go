package compress

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func Compress(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 1024 * 256)
	writer := bufio.NewWriterSize(w, 1024 * 256)
	for compressLine(reader, writer) {
	}
	writer.Flush()
}

func compressLine(r *bufio.Reader, w io.Writer) bool {
	name, idStrings, readErr := readLine(r)
	var packErr error
	if name != "" && len(idStrings) > 0 {
		packErr = Pack(w, []byte(name), CompressIds(idStrings))
	}

	if readErr == nil && packErr == nil {
		return true
	} else {
		log.Println("err read : ", readErr, " pack : ", packErr)
		return false
	}
}

func readLine(r *bufio.Reader) (name string, tagStrings []string, err error) {
	name, err = readString(r, '\t')
	if err != nil {
		return
	}
	tagString, err := readString(r, '\n')
	if tagString != "" {
		tagStrings = strings.Split(tagString, ",")
	}
	return
}

func readString(r *bufio.Reader, delimiter byte) (str string, err error) {
	str, err = r.ReadString(delimiter)
	if str != "" {
		str = str[:len(str)-1]
	}
	return
}
