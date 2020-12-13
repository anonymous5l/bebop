// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/200sc/bebop"
)

var _ bebop.Record = &Foo{}

type Foo struct {
	Bar Bar
}

func(bbp Foo) EncodeBebop(w io.Writer) (err error) {
	err = (bbp.Bar).EncodeBebop(w)
	if err != nil {
		return err
	}
	return nil
}

func(bbp *Foo) DecodeBebop(r io.Reader) (err error) {
	err = (&bbp.Bar).DecodeBebop(r)
	if err != nil {
		return err
	}
	return nil
}

func(bbp *Foo) bodyLen() (uint32) {
	bodyLen := uint32(0)
	bodyLen += (bbp.Bar).bodyLen()
	return bodyLen
}

var _ bebop.Record = &Bar{}

type Bar struct {
	X *float64
	Y *float64
	Z *float64
}

func(bbp Bar) EncodeBebop(w io.Writer) (err error) {
	binary.Write(w, binary.LittleEndian, bbp.bodyLen())
	if bbp.X != nil {
		w.Write([]byte{1})
		binary.Write(w, binary.LittleEndian, *bbp.X)
	}
	if bbp.Y != nil {
		w.Write([]byte{2})
		binary.Write(w, binary.LittleEndian, *bbp.Y)
	}
	if bbp.Z != nil {
		w.Write([]byte{3})
		binary.Write(w, binary.LittleEndian, *bbp.Z)
	}
	w.Write([]byte{0})
	return nil
}

func(bbp *Bar) DecodeBebop(ior io.Reader) (err error) {
	var bodyLen uint32
	var fieldNum byte
	binary.Read(ior, binary.LittleEndian, &bodyLen)
	body := make([]byte, bodyLen)
	ior.Read(body)
	r := bytes.NewReader(body)
	for r.Len() > 1 {
		fieldNum, _ = r.ReadByte()
		switch fieldNum {
		case 1:
			bbp.X = new(float64)
			binary.Read(r, binary.LittleEndian, bbp.X)
		case 2:
			bbp.Y = new(float64)
			binary.Read(r, binary.LittleEndian, bbp.Y)
		case 3:
			bbp.Z = new(float64)
			binary.Read(r, binary.LittleEndian, bbp.Z)
		default:
			return nil
		}
	}
	return nil
}

func(bbp *Bar) bodyLen() (uint32) {
	bodyLen := uint32(1)
	if bbp.X != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Y != nil {
		bodyLen += 1
		bodyLen += 8
	}
	if bbp.Z != nil {
		bodyLen += 1
		bodyLen += 8
	}
	return bodyLen
}
