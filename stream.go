package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

var wg sync.WaitGroup

const nbParts int = 6

type part struct {
	start int
	end   int
}

func main() {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	var parts []part

	end := 0
	start := 0

	length := len(bytes)
	fmt.Println(length)

	for i := 0; i < nbParts; i++ {
		end = start + (length / nbParts)
		if start >= length {
			break
		}

		if end >= length {
			end = length - 1
		}

		pt := part{
			start: start,
			end:   end,
		}

		parts = append(parts, pt)
		start = end + 1
	}

	result := make([]byte, length)
	fmt.Println(parts)

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

	fmt.Println(string(result))
}
