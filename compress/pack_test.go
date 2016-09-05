package compress

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	writer := new(bytes.Buffer)
	word := []byte{'t', 'e', 's', 't'}
	tags := []byte{5 + 128}
	expected := []byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't', 5 + 128}

	err := Pack(writer, word, tags)

	if !reflect.DeepEqual(expected, writer.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, writer.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestUnpack(t *testing.T) {
	input := bytes.NewBuffer([]byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't', 5 + 128})
	word := []byte{'t', 'e', 's', 't'}
	tags := []byte{5 + 128}

	actualWord, actualTags, err := Unpack(input)

	if !reflect.DeepEqual(word, actualWord) {
		t.Errorf("\nexpected: %v\nactual: %v", word, actualWord)
	}
	if !reflect.DeepEqual(tags, actualTags) {
		t.Errorf("\nexpected: %v\nactual: %v", tags, actualTags)
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestUnpackEOF(t *testing.T) {
	input := bytes.NewBuffer([]byte{})
	_, _, err := Unpack(input)
	if err.Error() != "EOF" {
		t.Errorf("\nexpected: %v\nactual: %v", "EOF", err)
	}
}
