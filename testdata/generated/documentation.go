// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

type DepE uint32

const (
	// Deprecated: X in DepE
	DepE_X DepE = 1
)

// Documented enum 
type DocE uint32

const (
	// Documented constant 
	DocE_X DocE = 1
	// Deprecated: Y in DocE
	DocE_Y DocE = 2
	// Deprecated, documented constant 
	// Deprecated: Z in DocE
	DocE_Z DocE = 3
)

var _ bebop.Record = &DocS{}

// Documented struct 
type DocS struct {
	// Documented field 
	X int32
}

func (bbp DocS) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteInt32(w, bbp.X)
	return w.Err
}

func (bbp *DocS) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	{
		bbp.X = iohelp.ReadInt32(r)
	}
	return r.Err
}

func (bbp *DocS) bodyLen() uint32 {
	bodyLen := uint32(0)
	bodyLen += 4
	return bodyLen
}

var _ bebop.Record = &DepM{}

type DepM struct {
	// Deprecated: x in DepM
	X *int32
}

func (bbp DepM) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	binary.Write(w, binary.LittleEndian, bbp.bodyLen())
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *DepM) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 0:
			return er.Err
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		default:
			return er.Err
		}
	}
	return er.Err
}

func (bbp *DepM) bodyLen() uint32 {
	bodyLen := uint32(1)
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

var _ bebop.Record = &DocM{}

// Documented message 
type DocM struct {
	// Documented field 
	X *int32
	// Deprecated: y in DocM
	Y *int32
	// Deprecated, documented field 
	// Deprecated: z in DocM
	Z *int32
}

func (bbp DocM) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	binary.Write(w, binary.LittleEndian, bbp.bodyLen())
	if bbp.X != nil {
		w.Write([]byte{1})
		iohelp.WriteInt32(w, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		iohelp.WriteInt32(w, *bbp.Y)
	}
	if bbp.Z != nil {
		w.Write([]byte{3})
		iohelp.WriteInt32(w, *bbp.Z)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *DocM) DecodeBebop(ior io.Reader) (err error) {
	er := iohelp.NewErrorReader(ior)
	bodyLen := iohelp.ReadUint32(er)
	body := make([]byte, bodyLen)
	er.Read(body)
	r := iohelp.NewErrorReader(bytes.NewReader(body))
	for {
		switch iohelp.ReadByte(r) {
		case 0:
			return er.Err
		case 1:
			bbp.X = new(int32)
			*bbp.X = iohelp.ReadInt32(r)
		case 2:
			bbp.Y = new(int32)
			*bbp.Y = iohelp.ReadInt32(r)
		case 3:
			bbp.Z = new(int32)
			*bbp.Z = iohelp.ReadInt32(r)
		default:
			return er.Err
		}
	}
	return er.Err
}

func (bbp *DocM) bodyLen() uint32 {
	bodyLen := uint32(1)
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.Z != nil {
		bodyLen += 1
		bodyLen += 4
	}
	return bodyLen
}

