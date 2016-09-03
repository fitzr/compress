package compress

import (
	"math"
	"reflect"
	"testing"
)

func TestEncodeOneByte(t *testing.T) {
	input := uint64(5)
	expected := []byte{5 + 128}

	actual := Encode(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestEncodeTwoByte(t *testing.T) {
	input := uint64(130)
	expected := []byte{1, 2 + 128}

	actual := Encode(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestEncodeInt64Max(t *testing.T) {
	input := uint64(math.MaxUint64)
	expected := []byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255}

	actual := Encode(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestEncodeZero(t *testing.T) {
	input := uint64(0)
	expected := []byte{128}

	actual := Encode(input)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestDecodeOneByte(t *testing.T) {
	input := []byte{5 + 128}
	expected := uint64(5)

	actual := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestDecodeTwoByte(t *testing.T) {
	input := []byte{1, 2 + 128}
	expected := uint64(130)

	actual := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestDecodeInt64Max(t *testing.T) {
	input := []byte{1, 127, 127, 127, 127, 127, 127, 127, 127, 255}
	expected := uint64(math.MaxUint64)

	actual := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}

func TestDecodeZero(t *testing.T) {
	input := []byte{128}
	expected := uint64(0)

	actual := Decode(input)

	if actual != expected {
		t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
	}
}
