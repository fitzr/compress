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

	Pack(writer, word, tags)

	if !reflect.DeepEqual(expected, writer.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, writer.Bytes())
	}
}

func TestUnpack(t *testing.T) {
	input := bytes.NewBuffer([]byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't', 5 + 128})
	word := []byte{'t', 'e', 's', 't'}
	tags := []byte{5 + 128}

	actualWord, actualTags := Unpack(input)

	if !reflect.DeepEqual(word, actualWord) {
		t.Errorf("\nexpected: %v\nactual: %v", word, actualWord)
	}
	if !reflect.DeepEqual(tags, actualTags) {
		t.Errorf("\nexpected: %v\nactual: %v", tags, actualTags)
	}
}

func TestUnpackPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Error("did not panic")
		}
		err, ok := r.(error)
		if !ok || err.Error() != "EOF" {
			t.Error("invalid panic reason")
		}
	}()
	input := bytes.NewBuffer([]byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't'})

	Unpack(input)
}
