package rnd

import (
	"crypto/rand"
	"unsafe"
)

func NewCryptoRandSeed() int64 {
	var seed [8]byte
	rand.Read(seed[:])
	return *(*int64)(unsafe.Pointer(&seed[0]))
}
