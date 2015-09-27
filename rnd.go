package rnd

import (
	"hash"
	"math/rand"
	"unsafe"
)

type (
	HashSource struct {
		h    hash.Hash
		seed []byte
	}
)

func NewHashSource(f func() hash.Hash, seed int64) rand.Source {
	s := &HashSource{h: f()}
	s.Seed(seed)
	return s
}

func (s *HashSource) Seed(seed int64) {
	seed8 := *(*[8]byte)(unsafe.Pointer(&seed))
	s.seed = seed8[:]
	s.h.Reset()
}

func (s *HashSource) Int63() int64 {
	s.hash()
	const mask = int64(^uint64(0) >> 1)
	return *(*int64)(unsafe.Pointer(&s.seed[0])) & mask
}

func (s *HashSource) hash() {
	s.h.Write(s.seed)
	copy(s.seed, s.h.Sum(nil))
}
