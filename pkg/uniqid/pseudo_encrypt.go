package uniqid

import (
	"math"
)

// pseudoEncrypt generate a integer id by val
// inspired by https://wiki.postgresql.org/wiki/Pseudo_encrypt
func pseudoEncrypt(val uint64) uint64 {
	var (
		l1 = uint64((val >> 32) & 0xffff)
		r1 = uint64(val & 0xffff)
		l2 uint64
		r2 uint64
	)
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ uint64(math.Round((float64((1366*r1+150889)%714025)/714025.0)*32767.0))
		l1 = l2
		r1 = r2
	}
	return uint64((r1 << 32) + l1)
}
