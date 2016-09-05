package compress

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func Compress(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	for compressLine(*reader, w) {
	}
}

func compressLine(r bufio.Reader, w io.Writer) bool {
	name, idStrings, err := readLine(r)
	if name != "" && len(idStrings) > 0 {
		Pack(w, []byte(name), CompressIds(idStrings))
	}

	if err == nil {
		return true
	} else {
		log.Println("read line error : ", err)
		return false
	}
}

func readLine(r bufio.Reader) (name string, tagStrings []string, err error) {
	name, err = readString(r, '\t')
	if err != nil {
		return
	}
	tagString, err := readString(r, '\n')
	if err != nil {
		return
	}
	tagStrings = strings.Split(tagString, ",")
	return
}

func readString(r bufio.Reader, delimiter byte) (str string, err error) {
	str, err = r.ReadString(delimiter)
	if err == nil {
		str = str[:len(str)-1]
	}
	return
}
