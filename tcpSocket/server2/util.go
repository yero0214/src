package main

import (
	"math/rand"
	"time"
	"unsafe"
)

func broadCast(buffer []byte) {
	for _, v := range users {
		v.Conn.Write(buffer)
	}
}

func randomNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func byteSliceToFloat32Slice(src []byte) []float32 {
	if len(src) == 0 {
		return nil
	}

	l := len(src) / 4
	ptr := unsafe.Pointer(&src[0])
	// It is important to keep in mind that the Go garbage collector
	// will not interact with this data, and that if src if freed,
	// the behavior of any Go code using the slice is nondeterministic.
	// Reference: https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	return (*[1 << 26]float32)((*[1 << 26]float32)(ptr))[:l:l]
}
