package rnd

import (
	"hash"
	"math/rand"
	"unsafe"
)

type (
	Source struct {
		h    hash.Hash
		seed []byte
	}
)

func NewHashSource(f func() hash.Hash, seed []byte) rand.Source {
	if len(seed) == 0 {
		seed = []byte{0}
	}
	return &Source{
		seed: seed,
		h:    f(),
	}
}

func (s *Source) Seed(seed int64) {
	seed8 := *(*[8]byte)(unsafe.Pointer(&seed))
	s.seed = seed8[:]
}

func (s *Source) Int63() int64 {
	s.hash()
	const mask = int64(^uint64(0) >> 1)
	return *(*int64)(unsafe.Pointer(&s.seed[0])) & mask
}

func (s *Source) hash() {
	for i := 0; i < len(s.seed); i++ {
		s.h.Write(s.seed)
		sum := s.h.Sum(nil)
		for j := 0; i < len(s.seed) && j < len(sum); i, j = i+1, j+1 {
			s.seed[i] = sum[j]
		}
	}
}
