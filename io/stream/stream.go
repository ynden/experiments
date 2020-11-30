package stream

import (
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

const nbParts int = 6

// Part is a struct that represents
// a chunk.
type Part struct {
	start int
	end   int
}

// Read takes a number of parts, and a file path as parameters.
// It retrieves a file, and split it as byte chunks of 'nbParts' parts.
// It returns the file's text.
func Read(nbParts int, filePath string) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	var parts []Part
	var start, end int
	length := len(bytes)

	for i := 0; i < nbParts; i++ {
		end = start + (length / nbParts)
		if start >= length {
			break
		}

		if end >= length {
			end = length - 1
		}

		pt := Part{
			start: start,
			end:   end,
		}

		parts = append(parts, pt)
		start = end + 1
	}

	result := make([]byte, length)

	for i := 0; i < len(parts); i++ {
		start := parts[i].start
		end := parts[i].end

		wg.Add(1)
		go func() {
			defer wg.Done()
			for idx, el := range bytes[start : end+1] {
				result[start+idx] = el // placing elements at the right position
			}
		}()

	}

	wg.Wait()

	return result
}
