//go:build go1.18 && !go1.20
// +build go1.18,!go1.20

package udecimal

import (
	"reflect"
	"unsafe"
)

func unsafeStringToBytes(s string) []byte {
	// return *(*[]byte)(unsafe.Pointer(&s))
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	var bytes []byte
	bytes = unsafe.Slice((*byte)(unsafe.Pointer(hdr.Data)), hdr.Len)
	return bytes
}
