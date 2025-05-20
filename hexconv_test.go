package hexconvgo

import (
	"testing"
)

func TestStrToHex(t *testing.T) {
	result := StrToHex("This is a test")
	expected := "54 68 69 73 20 69 73 20 61 20 74 65 73 74"

	if result != expected {
		t.Errorf("StrToHex(\"This is a test\") => %s; expected %s", result, expected)
	}
}
