// Code generated by bebopc-go; DO NOT EDIT.

package generated

import (
	"encoding/binary"
	"io"

	"github.com/200sc/bebop"
)

var _ bebop.Record = &BasicArrays{}

type BasicArrays struct {
	A_bool []bool
	A_byte []byte
	A_int16 []int16
	A_uint16 []uint16
	A_int32 []int32
	A_uint32 []uint32
	A_int64 []int64
	A_uint64 []uint64
	A_float32 []float32
	A_float64 []float64
	A_string []string
	A_guid [][16]byte
}

func(bbp BasicArrays) EncodeBebop(w io.Writer) (err error) {
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_bool)))
	for _, elem := range bbp.A_bool {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_byte)))
	for _, elem := range bbp.A_byte {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_int16)))
	for _, elem := range bbp.A_int16 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_uint16)))
	for _, elem := range bbp.A_uint16 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_int32)))
	for _, elem := range bbp.A_int32 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_uint32)))
	for _, elem := range bbp.A_uint32 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_int64)))
	for _, elem := range bbp.A_int64 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_uint64)))
	for _, elem := range bbp.A_uint64 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_float32)))
	for _, elem := range bbp.A_float32 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_float64)))
	for _, elem := range bbp.A_float64 {
		binary.Write(w, binary.LittleEndian, elem)
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_string)))
	for _, elem := range bbp.A_string {
		binary.Write(w, binary.LittleEndian, uint32(len(elem)))
		w.Write([]byte(elem))
	}
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A_guid)))
	for _, elem := range bbp.A_guid {
		w.Write(elem[:])
	}
	return nil
}

func(bbp *BasicArrays) DecodeBebop(r io.Reader) (err error) {
	var ln uint32
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(bool)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_bool = append(bbp.A_bool, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(byte)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_byte = append(bbp.A_byte, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(int16)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_int16 = append(bbp.A_int16, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(uint16)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_uint16 = append(bbp.A_uint16, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(int32)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_int32 = append(bbp.A_int32, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(uint32)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_uint32 = append(bbp.A_uint32, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(int64)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_int64 = append(bbp.A_int64, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(uint64)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_uint64 = append(bbp.A_uint64, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(float32)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_float32 = append(bbp.A_float32, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(float64)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A_float64 = append(bbp.A_float64, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(string)
		*elem1 = bebop.ReadString(r)
		bbp.A_string = append(bbp.A_string, *elem1)
	}
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new([16]byte)
		*elem1 = bebop.ReadGUID(r)
		bbp.A_guid = append(bbp.A_guid, *elem1)
	}
	return nil
}

func(bbp *BasicArrays) bodyLen() (uint32) {
	bodyLen := uint32(0)
	bodyLen += 4
	for _ = range bbp.A_bool {
		bodyLen += 1
	}
	bodyLen += 4
	for _ = range bbp.A_byte {
		bodyLen += 1
	}
	bodyLen += 4
	for _ = range bbp.A_int16 {
		bodyLen += 2
	}
	bodyLen += 4
	for _ = range bbp.A_uint16 {
		bodyLen += 2
	}
	bodyLen += 4
	for _ = range bbp.A_int32 {
		bodyLen += 4
	}
	bodyLen += 4
	for _ = range bbp.A_uint32 {
		bodyLen += 4
	}
	bodyLen += 4
	for _ = range bbp.A_int64 {
		bodyLen += 8
	}
	bodyLen += 4
	for _ = range bbp.A_uint64 {
		bodyLen += 8
	}
	bodyLen += 4
	for _ = range bbp.A_float32 {
		bodyLen += 4
	}
	bodyLen += 4
	for _ = range bbp.A_float64 {
		bodyLen += 8
	}
	bodyLen += 4
	for _, elem := range bbp.A_string {
		bodyLen += 4
		bodyLen += uint32(len(elem))
	}
	bodyLen += 4
	for _ = range bbp.A_guid {
		bodyLen += 16
	}
	return bodyLen
}

var _ bebop.Record = &TestInt32Array{}

type TestInt32Array struct {
	A []int32
}

func(bbp TestInt32Array) EncodeBebop(w io.Writer) (err error) {
	binary.Write(w, binary.LittleEndian, uint32(len(bbp.A)))
	for _, elem := range bbp.A {
		binary.Write(w, binary.LittleEndian, elem)
	}
	return nil
}

func(bbp *TestInt32Array) DecodeBebop(r io.Reader) (err error) {
	var ln uint32
	ln = uint32(0)
	binary.Read(r, binary.LittleEndian, &ln)
	for i := uint32(0); i < ln; i++ {
		elem1 := new(int32)
		binary.Read(r, binary.LittleEndian, elem1)
		bbp.A = append(bbp.A, *elem1)
	}
	return nil
}

func(bbp *TestInt32Array) bodyLen() (uint32) {
	bodyLen := uint32(0)
	bodyLen += 4
	for _ = range bbp.A {
		bodyLen += 4
	}
	return bodyLen
}
