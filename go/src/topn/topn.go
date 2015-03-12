// Implementation according to http://rd.springer.com/chapter/10.1007/3-540-45465-9_59#page-1
package topn

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"io"
	"sort"

	"gotl"
	"logging"
)

type FrequentItem struct {
	item  Item
	count int64
}

type FrequentItemSlice []FrequentItem

func (s FrequentItemSlice) Len() int           { return len(s) }
func (s FrequentItemSlice) Less(i, j int) bool { return s[i].count > s[j].count }
func (s FrequentItemSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type FrequentItemMonitor struct {
	TopItems     []FrequentItem
	MinFrequency int64
	TotalHits    int64

	// index in TopItems
	keyToIndex  map[string]int
	numTopItems int
	numHashes   int
	numBuckets  int
	buckets     [][]int64
	hash32      hash.Hash32
}

func NewFrequentItemMonitor(numTopItems int) *FrequentItemMonitor {
	return NewCustomizedFrequentItemMonitor(numTopItems, 20, 16, -3)
}

func NewCustomizedFrequentItemMonitor(numTopItems, logNumItems, logNumUniqueItems, logErrorRate int) *FrequentItemMonitor {
	numHashes := logNumItems
	numBuckets := 8 * (logNumUniqueItems - logErrorRate)
	buckets := make([][]int64, numHashes)
	for i := 0; i < numHashes; i++ {
		buckets[i] = make([]int64, numBuckets)
	}
	return &FrequentItemMonitor{
		[]FrequentItem{}, 0, 0,
		make(map[string]int), numTopItems, numHashes, numBuckets, buckets, fnv.New32a()}
}

func (m *FrequentItemMonitor) Observe(item Item) {
	m.ObserveN(item, 1)
}

func (m *FrequentItemMonitor) hashFunc(key string, seed int) (bucket, hashVal int64) {
	m.hash32.Reset()
	if err := binary.Write(m.hash32, binary.BigEndian, uint32(seed)); err != nil {
		logging.Fatalf("Failed to write seed: %s!", err)
	}
	if _, err := io.WriteString(m.hash32, key); err != nil {
		logging.Fatalf("Failed to write key: %s!", err)
	}
	bucket = int64(m.hash32.Sum32()) % int64(m.numBuckets)
	if err := binary.Write(m.hash32, binary.BigEndian, uint32(m.numHashes)); err != nil {
		logging.Fatalf("Failed to write #hash: %s!", err)
	}
	hashVal = 2*(int64(m.hash32.Sum32())%2) - 1
	return
}

func (m *FrequentItemMonitor) ObserveN(item Item, n int64) {
	m.TotalHits += n
	guesses := make([]int64, m.numHashes)
	for i := 0; i < m.numHashes; i++ {
		bucket, hashVal := m.hashFunc(item.Key(), i)
		m.buckets[i][bucket] += hashVal * n
		guesses[i] = m.buckets[i][bucket] * hashVal
		if guesses[i] < n {
			guesses[i] = n
		}
	}
	if index, found := m.keyToIndex[item.Key()]; found {
		m.TopItems[index].count += n
		return
	}

	gotl.Median(gotl.Int64Slice(guesses))
	observesGuess := guesses[len(guesses)/2]
	if len(m.TopItems) < m.numTopItems {
		m.keyToIndex[item.Key()] = len(m.TopItems)
		m.TopItems = append(m.TopItems, FrequentItem{item, observesGuess})
		if observesGuess < m.MinFrequency || m.MinFrequency == 0 {
			m.MinFrequency = observesGuess
		}
	} else if observesGuess > m.MinFrequency {
		toReplace := -1
		minFrequency := observesGuess
		m.MinFrequency = observesGuess
		for index, value := range m.TopItems {
			if value.count < minFrequency {
				m.MinFrequency = minFrequency
				toReplace = index
				minFrequency = value.count
			} else if value.count < m.MinFrequency {
				m.MinFrequency = value.count
			}
		}
		if toReplace != -1 {
			delete(m.keyToIndex, m.TopItems[toReplace].item.Key())
			m.TopItems[toReplace] = FrequentItem{item, observesGuess}
			m.keyToIndex[item.Key()] = toReplace
		}
	}
}

func (m *FrequentItemMonitor) SortedTopItems() []FrequentItem {
	sortedItems := []FrequentItem{}
	for _, value := range m.TopItems {
		sortedItems = append(sortedItems, value)
	}
	sort.Sort(FrequentItemSlice(sortedItems))
	return sortedItems
}
