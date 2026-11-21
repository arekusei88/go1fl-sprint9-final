package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates a slice of random integers.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return make([]int, 0)
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
	}

	return elements
}

// maximum returns the maximum value in data.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum value in data, searching its parts concurrently.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	maxValues := make([]int, CHUNKS)

	// Seed with data[0]: a chunk left empty when len(data) < CHUNKS keeps it
	// instead of a zero, which would falsely win when all data is negative.
	for i := range maxValues {
		maxValues[i] = data[0]
	}

	var wg sync.WaitGroup
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			// the last chunk takes the remainder so that elements are not lost
			end = len(data)
		}
		if start == end {
			continue
		}

		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
		}(data[start:end], i)
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxValueSingle := maximum(elements)
	elapsedSingle := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n", maxValueSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxValueConcurrent := maxChunks(elements)
	elapsedConcurrent := time.Since(startTime).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d mcs\n", maxValueConcurrent, elapsedConcurrent)
}
