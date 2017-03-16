package depfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataBufferSetBytesGetBytes(t *testing.T) {
	deps := CreateDepfileBuilder("a b")
	deps.Add("c d\\e")
	deps.Add("f")
	assert.Equal(t, "a\\ b: c\\ d\\\\e f\n", deps.String())
}
