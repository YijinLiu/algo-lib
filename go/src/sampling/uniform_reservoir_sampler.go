package sampling

import (
	"math/rand"
	"time"

	"logging"
)

type ReservoirSampler struct {
	common  *ReservoirSamplerCommon
	randGen *rand.Rand
}

func NewReservoirSampler(numSamples int64) *ReservoirSampler {
	return &ReservoirSampler{NewReservoirSamplerCommon(numSamples), rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewReservoirSamplerWithRand(numSamples int64, randGen *rand.Rand) *ReservoirSampler {
	return &ReservoirSampler{NewReservoirSamplerCommon(numSamples), randGen}
}

func (s *ReservoirSampler) NumItems() int64 {
	return s.common.NumItems
}

func (s *ReservoirSampler) Add(item Item) {
	v := s.randGen.Float64()
	s.common.AddWithV(item, v)
}

func (s *ReservoirSampler) Samples() []Item {
	return s.common.samples
}

type HashBasedReservoirSampler struct {
	numItems   int64
	numSamples int64
	samples    []Item
}

func NewHashBasedReservoirSampler(numSamples int64) *HashBasedReservoirSampler {
	if numSamples <= 0 {
		logging.Fatalf("numSamples must > 0: %d!", numSamples)
	}
	return &HashBasedReservoirSampler{0, numSamples, []Item{}}
}

func (s *HashBasedReservoirSampler) NumItems() int64 {
	return s.numItems
}

// hashV could also be some random number.
func (s *HashBasedReservoirSampler) AddWithHash(item Item, hashV uint64) {
	s.numItems++
	if int64(len(s.samples)) < s.numSamples {
		s.samples = append(s.samples, item)
		return
	}

	// >> 1 to avoid negative values.
	index := int64(hashV>>1) % s.numItems
	if index < s.numSamples {
		s.samples[index] = item
	}
}

func (s *HashBasedReservoirSampler) Samples() []Item {
	return s.samples
}

// Implementation based on http://www.cs.umd.edu/~samir/498/vitter.pdf
type FastForwardReservoirSampler struct {
	numItems   int64
	numSamples int64
	samples    []Item
	randGen    *rand.Rand
}

func NewFastForwardReservoirSampler(numSamples int64) *FastForwardReservoirSampler {
	return NewFastForwardReservoirSamplerWithRand(numSamples, rand.New(rand.NewSource(time.Now().UnixNano())))
}

func NewFastForwardReservoirSamplerWithRand(numSamples int64, randGen *rand.Rand) *FastForwardReservoirSampler {
	return &FastForwardReservoirSampler{0, numSamples, []Item{}, randGen}
}

func (s *FastForwardReservoirSampler) NumItems() int64 {
	return s.numItems
}

// Returns negative value means skip to the end.
func (s *FastForwardReservoirSampler) GetSkips() int64 {
	if s.numItems < s.numSamples {
		return 0
	}

	const MINIMAL_FASTFORWARD_CHANCE = 1e-20
	chance := s.randGen.Float64()
	if chance < chance {
		return -1
	}

	var currentChance float64 = 1
	var skips int64 = 0
	for {
		currentChance *= float64(s.numItems+skips+1-s.numSamples) / float64(s.numItems+skips+1)
		if currentChance <= chance {
			return skips
		}

	}
}

func (s *FastForwardReservoirSampler) Add(item Item) {
	s.numItems++
	if s.numItems <= s.numSamples {
		s.samples = append(s.samples, item)
		return
	}
	index := s.randGen.Int63() % s.numSamples
	s.samples[index] = item
}

func (s *FastForwardReservoirSampler) Samples() []Item {
	return s.samples
}
