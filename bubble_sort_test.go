package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}

	expected := []int{3, 4, 7, 9, 42, 69, 420}

	assert.Equal(t, expected, bubble_sort(arr))
}
