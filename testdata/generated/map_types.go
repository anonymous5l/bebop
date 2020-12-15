// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/200sc/bebop"
	"github.com/200sc/bebop/iohelp"
)

var _ bebop.Record = &S{}

type S struct {
	x int32
	y int32
}

func (bbp S) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteInt32(w, bbp.x)
	iohelp.WriteInt32(w, bbp.y)
	return w.Err
}

func (bbp *S) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	{
		bbp.x = iohelp.ReadInt32(r)
	}
	{
		bbp.y = iohelp.ReadInt32(r)
	}
	return r.Err
}

func (bbp *S) bodyLen() uint32 {
	bodyLen := uint32(0)
	bodyLen += 4
	bodyLen += 4
	return bodyLen
}

func (bbp S) GetX() int32 {
	return bbp.x
}

func (bbp S) GetY() int32 {
	return bbp.y
}

var _ bebop.Record = &SomeMaps{}

type SomeMaps struct {
	M1 map[bool]bool
	M2 map[string]map[string]string
	M3 []map[int32][]map[bool]S
	M4 []map[string][]float32
	M5 map[[16]byte]M
}

func (bbp SomeMaps) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	iohelp.WriteUint32(w, uint32(len(bbp.M1)))
	for k, v := range bbp.M1 {
		iohelp.WriteBool(w, k)
		iohelp.WriteBool(w, v)
	}
	iohelp.WriteUint32(w, uint32(len(bbp.M2)))
	for k, v := range bbp.M2 {
		iohelp.WriteUint32(w, uint32(len(k)))
		w.Write([]byte(k))
		iohelp.WriteUint32(w, uint32(len(v)))
		for k, v := range v {
			iohelp.WriteUint32(w, uint32(len(k)))
			w.Write([]byte(k))
			iohelp.WriteUint32(w, uint32(len(v)))
			w.Write([]byte(v))
		}
	}
	iohelp.WriteUint32(w, uint32(len(bbp.M3)))
	for _, elem := range bbp.M3 {
		iohelp.WriteUint32(w, uint32(len(elem)))
		for k, v := range elem {
			iohelp.WriteInt32(w, k)
			iohelp.WriteUint32(w, uint32(len(v)))
			for _, elem := range v {
				iohelp.WriteUint32(w, uint32(len(elem)))
				for k, v := range elem {
					iohelp.WriteBool(w, k)
					err = (v).EncodeBebop(w)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	iohelp.WriteUint32(w, uint32(len(bbp.M4)))
	for _, elem := range bbp.M4 {
		iohelp.WriteUint32(w, uint32(len(elem)))
		for k, v := range elem {
			iohelp.WriteUint32(w, uint32(len(k)))
			w.Write([]byte(k))
			iohelp.WriteUint32(w, uint32(len(v)))
			for _, elem := range v {
				iohelp.WriteFloat32(w, elem)
			}
		}
	}
	iohelp.WriteUint32(w, uint32(len(bbp.M5)))
	for k, v := range bbp.M5 {
		iohelp.WriteGUID(w, k)
		err = (v).EncodeBebop(w)
		if err != nil {
			return err
		}
	}
	return w.Err
}

func (bbp *SomeMaps) DecodeBebop(ior io.Reader) (err error) {
	r := iohelp.NewErrorReader(ior)
	{
		ln2 := iohelp.ReadUint32(r)
		bbp.M1 = make(map[bool]bool)
		for i := uint32(0); i < ln2; i++ {
			k := new(bool)
			*k = iohelp.ReadBool(r)
			elem2 := new(bool)
			*elem2 = iohelp.ReadBool(r)
			(bbp.M1)[*k] = *elem2
		}
	}
	{
		ln2 := iohelp.ReadUint32(r)
		bbp.M2 = make(map[string]map[string]string)
		for i := uint32(0); i < ln2; i++ {
			k := new(string)
			*k = iohelp.ReadString(r)
			elem2 := new(map[string]string)
			ln3 := iohelp.ReadUint32(r)
			*elem2 = make(map[string]string)
			for i := uint32(0); i < ln3; i++ {
				k := new(string)
				*k = iohelp.ReadString(r)
				elem3 := new(string)
				*elem3 = iohelp.ReadString(r)
				(*elem2)[*k] = *elem3
			}
			(bbp.M2)[*k] = *elem2
		}
	}
	{
		ln2 := iohelp.ReadUint32(r)
		for i := uint32(0); i < ln2; i++ {
			elem2 := new(map[int32][]map[bool]S)
			ln3 := iohelp.ReadUint32(r)
			*elem2 = make(map[int32][]map[bool]S)
			for i := uint32(0); i < ln3; i++ {
				k := new(int32)
				*k = iohelp.ReadInt32(r)
				elem3 := new([]map[bool]S)
				ln4 := iohelp.ReadUint32(r)
				for i := uint32(0); i < ln4; i++ {
					elem4 := new(map[bool]S)
					ln5 := iohelp.ReadUint32(r)
					*elem4 = make(map[bool]S)
					for i := uint32(0); i < ln5; i++ {
						k := new(bool)
						*k = iohelp.ReadBool(r)
						elem5 := new(S)
						err = (elem5).DecodeBebop(r)
						if err != nil {
							return err
						}
						(*elem4)[*k] = *elem5
					}
					*elem3 = append(*elem3, *elem4)
				}
				(*elem2)[*k] = *elem3
			}
			bbp.M3 = append(bbp.M3, *elem2)
		}
	}
	{
		ln2 := iohelp.ReadUint32(r)
		for i := uint32(0); i < ln2; i++ {
			elem2 := new(map[string][]float32)
			ln3 := iohelp.ReadUint32(r)
			*elem2 = make(map[string][]float32)
			for i := uint32(0); i < ln3; i++ {
				k := new(string)
				*k = iohelp.ReadString(r)
				elem3 := new([]float32)
				ln4 := iohelp.ReadUint32(r)
				for i := uint32(0); i < ln4; i++ {
					elem4 := new(float32)
					*elem4 = iohelp.ReadFloat32(r)
					*elem3 = append(*elem3, *elem4)
				}
				(*elem2)[*k] = *elem3
			}
			bbp.M4 = append(bbp.M4, *elem2)
		}
	}
	{
		ln2 := iohelp.ReadUint32(r)
		bbp.M5 = make(map[[16]byte]M)
		for i := uint32(0); i < ln2; i++ {
			k := new([16]byte)
			*k = iohelp.ReadGUID(r)
			elem2 := new(M)
			err = (elem2).DecodeBebop(r)
			if err != nil {
				return err
			}
			(bbp.M5)[*k] = *elem2
		}
	}
	return r.Err
}

func (bbp *SomeMaps) bodyLen() uint32 {
	bodyLen := uint32(0)
	bodyLen += 4
	for range bbp.M1 {
		bodyLen += 1
		bodyLen += 1
	}
	bodyLen += 4
	for k, v := range bbp.M2 {
		bodyLen += 4
		bodyLen += uint32(len(k))
		bodyLen += 4
		for k, v := range v {
			bodyLen += 4
			bodyLen += uint32(len(k))
			bodyLen += 4
			bodyLen += uint32(len(v))
		}
	}
	bodyLen += 4
	for _, elem := range bbp.M3 {
		bodyLen += 4
		for _, v := range elem {
			bodyLen += 4
			bodyLen += 4
			for _, elem := range v {
				bodyLen += 4
				for _, v := range elem {
					bodyLen += 1
					bodyLen += (v).bodyLen()
				}
			}
		}
	}
	bodyLen += 4
	for _, elem := range bbp.M4 {
		bodyLen += 4
		for k, v := range elem {
			bodyLen += 4
			bodyLen += uint32(len(k))
			bodyLen += 4
			for range v {
				bodyLen += 4
			}
		}
	}
	bodyLen += 4
	for _, v := range bbp.M5 {
		bodyLen += 16
		bodyLen += (v).bodyLen()
	}
	return bodyLen
}

var _ bebop.Record = &M{}

type M struct {
	A *float32
	B *float64
}

func (bbp M) EncodeBebop(iow io.Writer) (err error) {
	w := iohelp.NewErrorWriter(iow)
	binary.Write(w, binary.LittleEndian, bbp.bodyLen())
	if bbp.A != nil {
		w.Write([]byte{1})
		iohelp.WriteFloat32(w, *bbp.A)
	}
	if bbp.B != nil {
		w.Write([]byte{2})
		iohelp.WriteFloat64(w, *bbp.B)
	}
	w.Write([]byte{0})
	return w.Err
}

func (bbp *M) DecodeBebop(ior io.Reader) (err error) {
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
			bbp.A = new(float32)
			*bbp.A = iohelp.ReadFloat32(r)
		case 2:
			bbp.B = new(float64)
			*bbp.B = iohelp.ReadFloat64(r)
		default:
			return er.Err
		}
	}
	return er.Err
}

func (bbp *M) bodyLen() uint32 {
	bodyLen := uint32(1)
	if bbp.A != nil {
		bodyLen += 1
		bodyLen += 4
	}
	if bbp.B != nil {
		bodyLen += 1
		bodyLen += 8
	}
	return bodyLen
}

