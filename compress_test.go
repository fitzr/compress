package compress

import (
	"bufio"
	"bytes"
	"testing"
)

func TestReadLine(t *testing.T) {
	text := `hpmini110	17064731
①Windows	1409171,14711245,18265928,21590872`
	input := bufio.NewReader(bytes.NewReader([]byte(text)))
	name, ids, err := readLine(input)

	if name != "hpmini110" {
		t.Errorf("\nexpected: %v\nactual: %v", "hpmini110", name)
	}
	if ids[0] != "17064731" {
		t.Errorf("\nexpected: %v\nactual: %v", "17064731", ids[0])
	}
	if err != nil {
		t.Error("error")
	}

	name, ids, err = readLine(input)
	if len(ids) != 4 {
		t.Errorf("\nexpected: %v\nactual: %v", 4, len(ids))
	}
}
