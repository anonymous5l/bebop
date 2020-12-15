package generated_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/200sc/bebop/testdata/generated"
)

var benchArray = &generated.BasicArrays{
	A_bool: []bool{
		true, false, true, false,
	},
	A_byte: []byte{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	},
	A_int16: []int16{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	},
	A_uint16: []uint16{
		0, 1, 2, 3, 4, 5, 6, 7, 8,
	},
	A_int32: []int32{
		0, 1, 234436345, 3, 4, 5, 634, 7, 8,
	},
	A_uint32: []uint32{
		0, 1, 2, 33453566, 4, 5, 634634, 7, 8,
	},
	A_int64: []int64{
		3436453450, 346345346531, 3463453452, 3, 4, 5346345345, 34634566, 7, 8,
	},
	A_uint64: []uint64{
		0, 1, 2, 3, 34634563454, 5, 6334534634, 7, 8,
	},
	A_float32: []float32{
		0, 341, 2, 34563453, 4, 5, 6, 7, 8,
	},
	A_float64: []float64{
		0, 1, 2, 345343, 3453464, 3453453635, 353453456, 7, 8555555555,
	},
	A_string: []string{
		"0123151234123123", "11234125123415124", "223412512341512341254", "31245123151234125123413", "1231251231512315124", "124123151234151234125", "61231512341541234123", "12315123412512341257", "81231451241234151234151",
	},
}

var benchArray2 *generated.BasicArrays

func BenchmarkMarshalBasicArrays(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		benchArray.EncodeBebop(w)
	}
}

func BenchmarkMarshalUnmarshalBasicArrays(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		benchArray.EncodeBebop(w)
		benchArray2 = &generated.BasicArrays{}
		benchArray2.DecodeBebop(w)
	}
}

func BenchmarkUnmarshalBasicArrays(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		benchArray.EncodeBebop(w)
		b.StartTimer()
		benchArray2 = &generated.BasicArrays{}
		benchArray2.DecodeBebop(w)
	}
}

func BenchmarkMarshalBasicArraysJSON(b *testing.B) {
	var w = &bytes.Buffer{}
	encoder := json.NewEncoder(w)
	for i := 0; i < b.N; i++ {
		encoder.Encode(benchArray)
	}
}

func BenchmarkMarshalUnmarshalBasicArraysJSON(b *testing.B) {
	var w = &bytes.Buffer{}
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(w)
	for i := 0; i < b.N; i++ {
		err := encoder.Encode(benchArray)
		if err != nil {
			panic(err)
		}
		decoder.Decode(&benchArray2)
	}
}

func BenchmarkUnmarshalBasicArraysJSON(b *testing.B) {
	var w = &bytes.Buffer{}
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(w)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		err := encoder.Encode(benchArray)
		if err != nil {
			panic(err)
		}
		b.StartTimer()
		decoder.Decode(&benchArray2)
	}
}

var benchMaps = &generated.SomeMaps{
	M1: map[bool]bool{
		true:  true,
		false: false,
	},
	M2: map[string]map[string]string{
		"hello": map[string]string{
			"world": "!",
		},
		"foo": map[string]string{
			"bar": "bizz",
		},
		"ursula": map[string]string{
			"k": "leguin",
		},
		"mario": map[string]string{
			"mario":    "",
			"luigi":    "",
			"brothers": "",
		},
	},
	M3: []map[int32][]map[bool]generated.S{
		{
			0: []map[bool]generated.S{{
				true:  generated.S{},
				false: generated.S{},
			}},
		}, {
			1: []map[bool]generated.S{{
				true:  generated.S{},
				false: generated.S{},
			}},
			2: []map[bool]generated.S{{
				true:  generated.S{},
				false: generated.S{},
			}},
			3: []map[bool]generated.S{{
				true:  generated.S{},
				false: generated.S{},
			}},
		}, {
			41111: []map[bool]generated.S{{
				true:  generated.S{},
				false: generated.S{},
			}},
		},
	},
	M4: []map[string][]float32{
		{
			"a": []float32{1321, 1423, 1423, 540, 12314, 1231, 4123, 1412, 1230, 4123, 123},
		},
	},
	M5: map[[16]byte]generated.M{
		[16]byte{5: 3}: generated.M{B: float64p(0.0000002)},
	},
}

var benchMaps2 *generated.SomeMaps

func BenchmarkMarshalSomeMap(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		benchMaps.EncodeBebop(w)
	}
}

func BenchmarkMarshalUnmarshalSomeMap(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		benchMaps.EncodeBebop(w)
		benchMaps2 = &generated.SomeMaps{}
		benchMaps2.DecodeBebop(w)
	}
}

func BenchmarkUnmarshalSomeMap(b *testing.B) {
	var w = &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		benchMaps.EncodeBebop(w)
		b.StartTimer()
		benchMaps2 = &generated.SomeMaps{}
		benchMaps2.DecodeBebop(w)
	}
}

// SomeMap cannot be parsed by json as its structure is unsupported
