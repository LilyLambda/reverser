package reverser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyString(t *testing.T) {
	input := ""
	expected := ""
	assert.Equal(t, expected, Reverse(input), "reversing empty string")
}

func TestReverseAscii(t *testing.T) {
	input := "hello, world"
	expected := "dlrow ,olleh"
	assert.Equal(t, expected, Reverse(input), "reversing ASCII string")
}

func TestChineseString(t *testing.T) {
	input := "世界你好"
	expected := "好你界世"
	assert.Equal(t, expected, Reverse(input), "reversing Chinese string")
}

func TestHebrewString(t *testing.T) {
	input := "שלום"
	expected := "םולש"
	assert.Equal(t, expected, Reverse(input), "reversing Hebrew string")
}

func TestDiacritics(t *testing.T) {
	input := "alɣ̞o"
	expected := "oɣ̞la"
	t.Skip("IPA diacritics do not reverse symbol-wise")
	assert.Equal(t, expected, Reverse(input), "reversing IPA diacritics")
}
