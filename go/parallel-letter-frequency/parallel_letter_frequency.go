package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {

	frequency := FreqMap{}
	var wg sync.WaitGroup

	ch := make(chan FreqMap)
	wg.Add(len(texts))

	for _, text := range texts {
		go func(text string, wg *sync.WaitGroup) {
			defer wg.Done()
			ch <- Frequency(text)
		}(text, &wg)
	}

	go func() {
		for {
			select {
			case fmap := <-ch:
				for k, v := range fmap {
					frequency[k] = frequency[k] + v
				}
			}
		}
	}()

	wg.Wait()
	close(ch)

	return frequency

}
