package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "zero size", size: 0, wantLen: 0},
		{name: "negative size", size: -10, wantLen: 0},
		{name: "single element", size: 1, wantLen: 1},
		{name: "regular size", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Errorf("generateRandomElements(%d) length = %d, want %d", tt.size, len(got), tt.wantLen)
			}
			for i, value := range got {
				if value < 0 {
					t.Errorf("generateRandomElements(%d)[%d] = %d, want a non-negative value", tt.size, i, value)
					break
				}
			}
		})
	}
}

func TestGenerateRandomElementsDifferent(t *testing.T) {
	first := generateRandomElements(100)
	second := generateRandomElements(100)

	same := true
	for i := range first {
		if first[i] != second[i] {
			same = false
			break
		}
	}
	if same {
		t.Error("two generateRandomElements(100) calls returned identical slices, want different values")
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "maximum in the middle", data: []int{3, 9, 1, 7, 5}, want: 9},
		{name: "maximum at the start", data: []int{100, 1, 2, 3}, want: 100},
		{name: "maximum at the end", data: []int{1, 2, 3, 100}, want: 100},
		{name: "duplicate maximums", data: []int{4, 8, 8, 2, 8}, want: 8},
		{name: "negative values", data: []int{-5, -1, -9}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "nil slice", data: nil, want: 0},
		{name: "single element", data: []int{42}, want: 42},
		{name: "fewer elements than chunks", data: []int{3, 9, 1}, want: 9},
		{name: "exactly one chunk size", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "length is not a multiple of chunks", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, want: 10},
		{name: "maximum in the remainder of the last chunk", data: []int{10, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 99}, want: 99},
		{name: "duplicate maximums", data: []int{4, 8, 8, 2, 8, 1, 1, 1, 1}, want: 8},
		{name: "negative values", data: []int{-5, -1, -9}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{0, 1, 7, 8, 9, 100, 1001} {
		data := generateRandomElements(size)
		if got, want := maxChunks(data), maximum(data); got != want {
			t.Errorf("maxChunks(generateRandomElements(%d)) = %d, want %d", size, got, want)
		}
	}
}
