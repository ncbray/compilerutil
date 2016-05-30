package fs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataBufferSetBytesGetBytes(t *testing.T) {
	payload := []byte{0, 1, 2, 3, 4, 5, 6}
	dbuf := &DataBuffer{}
	err := dbuf.SetBytes(payload)
	assert.Equal(t, nil, err)
	result, err := dbuf.GetBytes()
	assert.Equal(t, nil, err)
	assert.Equal(t, payload, result)
}

func TestDataBufferWriterReader(t *testing.T) {
	payload := []byte{0, 1, 2, 3, 4, 5, 6}
	dbuf := &DataBuffer{}
	w, err := dbuf.GetWriter()
	assert.Equal(t, nil, err)
	n, err := w.Write(payload)
	assert.Equal(t, nil, err)
	assert.Equal(t, 7, n)

	r, err := dbuf.GetReader()
	assert.Equal(t, nil, err)

	result := make([]byte, 7)
	n, err = r.Read(result)
	assert.Equal(t, nil, err)
	assert.Equal(t, 7, n)
	assert.Equal(t, payload, result)
}
