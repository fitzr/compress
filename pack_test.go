package compress

import (
	"bytes"
	"reflect"
	"testing"
)

func TestPack(t *testing.T) {
	writer := new(bytes.Buffer)
	name := []byte{'t', 'e', 's', 't'}
	ids := []byte{5 + 128}
	expected := []byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't', 5 + 128}

	err := Pack(writer, name, ids)

	if !reflect.DeepEqual(expected, writer.Bytes()) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, writer.Bytes())
	}
	if err != nil {
		t.Errorf("\nerror: %v", err)
	}
}

func TestUnpack(t *testing.T) {
	input := bytes.NewBuffer([]byte{4, 0, 0, 0, 1, 0, 0, 0, 't', 'e', 's', 't', 5 + 128})
	name := []byte{'t', 'e', 's', 't'}
	ids := []byte{5 + 128}

	actualName, actualIds, err := Unpack(input)

	if !reflect.DeepEqual(name, actualName) {
		t.Errorf("\nexpected: %v\nactual: %v", name, actualName)
	}
	if !reflect.DeepEqual(ids, actualIds) {
		t.Errorf("\nexpected: %v\nactual: %v", ids, actualIds)
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
