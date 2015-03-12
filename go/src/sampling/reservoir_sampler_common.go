package sampling

import (
	"logging"
)

type ReservoirSamplerCommon struct {
	NumItems int64

	numSamples int64
	samples    []Item
	sampleVs   []float64
	maxV       float64
}

func NewReservoirSamplerCommon(numSamples int64) *ReservoirSamplerCommon {
	if numSamples <= 0 {
		logging.Fatalf("numSamples must > 0: %d!", numSamples)
	}
	return &ReservoirSamplerCommon{0, numSamples, []Item{}, []float64{}, 0}
}

// Keep the N smallest Vs.
func (s *ReservoirSamplerCommon) AddWithV(item Item, v float64) {
	s.NumItems++

	if int64(len(s.samples)) < s.numSamples {
		s.samples = append(s.samples, item)
		s.sampleVs = append(s.sampleVs, v)
		if len(s.samples) == 1 || v > s.maxV {
			s.maxV = v
		}
		return
	}

	if v < s.maxV {
		toReplace := -1
		maxV := v
		s.maxV = v
		for index, value := range s.sampleVs {
			if value > maxV {
				s.maxV = maxV
				toReplace = index
				maxV = value
			} else if value > s.maxV {
				s.maxV = value
			}
		}
		s.samples[toReplace] = item
		s.sampleVs[toReplace] = v
	}
}
