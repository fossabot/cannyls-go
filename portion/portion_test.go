package portion

import (
	"github.com/stretchr/testify/assert"
	"github.com/thesues/cannyls-go/address"
	"testing"
)

func TestFreePortion(t *testing.T) {
	p := New(address.AddressFromU32(100), 50)
	assert.Equal(t, address.AddressFromU32(100), p.Start())
	assert.Equal(t, address.AddressFromU32(150), p.End())
	assert.Equal(t, uint32(50), p.Len())
}

func TestFreePortionAlloca(t *testing.T) {
	p := New(address.AddressFromU32(100), 150)
	p, alloc := p.SlicePart(30)
	assert.Equal(t, address.AddressFromU32(100), alloc.start)
	assert.Equal(t, uint16(30), alloc.len)

	assert.Equal(t, address.AddressFromU32(130), p.Start())
	assert.Equal(t, uint32(120), p.Len())
	assert.Equal(t, address.AddressFromU32(250), p.End())

	p, alloc = p.SlicePart(120)
	assert.Equal(t, address.AddressFromU32(130), alloc.start)
	assert.Equal(t, uint16(120), alloc.len)
	assert.Equal(t, uint32(0), p.Len())
	//assert.Equal(t, address.AddressFromU32(250), p.Start()

}