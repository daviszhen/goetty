package simple

import (
	"testing"

	"github.com/fagongzi/goetty/v2/buf"
	"github.com/stretchr/testify/assert"
)

func TestStringCodec(t *testing.T) {
	v := "hello world"
	buf := buf.NewByteBuf(32)
	codec := NewStringCodec()
	assert.NoError(t, codec.Encode(v, buf, nil), "TestStringCodec failed")
	readed, completed, err := codec.Decode(buf)
	assert.NoError(t, err, "TestStringCodec failed")
	assert.True(t, completed, "TestStringCodec failed")
	assert.Equal(t, v, readed, "TestStringCodec failed")
}
