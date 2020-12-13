// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"encoding/binary"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &ArrayOfStrings{}

type ArrayOfStrings struct {
	Strings []string
}

func (bbp ArrayOfStrings) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.ErrorWriter{Writer:iow}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.Strings)))
	for _, elem := range bbp.Strings {
		binary.Write(w, binary.LittleEndian, uint32(len(elem)))
		w.Write([]byte(elem))
	}
	return w.Err
}

func (bbp *ArrayOfStrings) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.ErrorReader{Reader:ior}
	var ln uint32
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(string)
		*elem1 = iohelp.ReadString(r)
		bbp.Strings = append(bbp.Strings, *elem1)
	}
	return r.Err
}

func (bbp *ArrayOfStrings) bodyLen() (uint32) {
	bodyLen := uint32(0)
	bodyLen += 4
	for _, elem := range bbp.Strings {
		bodyLen += 4
		bodyLen += uint32(len(elem))
	}
	return bodyLen
}

