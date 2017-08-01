// +build !fast

package reverser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
)

func TestHello(t *testing.T) {
	input := "hello"
	expected := "olleh"
	time.Sleep(25 * time.Second)
	assert.Equal(t, expected, Reverse(input), "reversing empty string")
}
