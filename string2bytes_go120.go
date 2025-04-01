//go:build go1.20
// +build go1.20

package udecimal

import "unsafe"

func unsafeStringToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
