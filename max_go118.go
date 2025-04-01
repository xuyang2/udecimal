//go:build go1.18 && !go1.21
// +build go1.18,!go1.21

package udecimal

func _max(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}
