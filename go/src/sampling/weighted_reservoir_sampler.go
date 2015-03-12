// Implementation based on http://doi.acm.org/10.1145/1141885.1141891

package sampling

import (
	"math/rand"
	"time"

	"logging"
)

type WeightedReservoirSampler struct {
	common  *ReservoirSamplerCommon
	randGen *rand.Rand
}

func NewWeightedReservoirSampler(numSamples int64) *WeightedReservoirSampler {
	return &WeightedReservoirSampler{NewReservoirSamplerCommon(numSamples), rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func NewWeightedReservoirSamplerWithRand(numSamples int64, randGen *rand.Rand) *WeightedReservoirSampler {
	return &WeightedReservoirSampler{NewReservoirSamplerCommon(numSamples), randGen}
}

func (s *WeightedReservoirSampler) NumItems() int64 {
	return s.common.NumItems
}

func (s *WeightedReservoirSampler) AddWithWeight(item Item, weight float64) {
	if weight <= 0.0 {
		logging.Fatalf("Weight must be positive number! (%f)", weight)
	}
	v := s.randGen.ExpFloat64() / weight
	s.common.AddWithV(item, v)
}

func (s *WeightedReservoirSampler) Samples() []Item {
	return s.common.samples
}
