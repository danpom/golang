package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	result := make(FreqMap)
	num := len(l)
	ch := make(chan FreqMap, num)

	var wg sync.WaitGroup

	wg.Add(num)

	//spawn a goroutine to check the frequency of each rune in each text of l
	for _, text := range l {
		go func(s string) {
			ch <- Frequency(s)
			wg.Done()
		}(text)
	}

	//wait for go routines to finish then close the channel
	go func() {
		wg.Wait()
		close(ch)
	}()

	//combine frequency from each channel
	for freqmap := range ch {
		for letter, freq := range freqmap {
			result[letter] += freq
		}
	}

	return result
}
