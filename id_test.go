package compress

import (
	"reflect"
	"testing"
)

func TestCompressIds(t *testing.T) {
	input := []string{"11886295", "15970919", "15970932"}
	expected := []byte{5, 85, 61, 215, 1, 121, 39, 144, 141}

	actual := CompressIds(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestDecompressIds(t *testing.T) {
	input := []byte{5, 85, 61, 215, 1, 121, 39, 144, 141}
	expected := []string{"11886295", "15970919", "15970932"}

	actual := DecompressIds(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
