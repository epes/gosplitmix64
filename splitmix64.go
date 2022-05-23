package gosplitmix64

import "math/rand"

// SplitMix64 is a go translation of https://prng.di.unimi.it/splitmix64.c
// Implements rand.Source and rand.Source64
type SplitMix64 struct {
	state uint64
}

func (s *SplitMix64) Seed(seed int64) {
	s.state = uint64(seed)
}

func (s *SplitMix64) Next() uint64 {
	s.state += 0x9e3779b97f4a7c15
	result := s.state
	result = (result ^ (result >> 30)) * 0xbf58476d1ce4e5b9
	result = (result ^ (result >> 27)) * 0x94d049bb133111eb
	result = result ^ (result >> 31)
	return result
}

func (s *SplitMix64) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *SplitMix64) Uint64() uint64 {
	return s.Next()
}

func New(seed int64) *SplitMix64 {
	s := &SplitMix64{}
	s.Seed(seed)
	return s
}

func NewSource() rand.Source {
	return &SplitMix64{}
}

func NewSource64() rand.Source64 {
	return &SplitMix64{}
}
