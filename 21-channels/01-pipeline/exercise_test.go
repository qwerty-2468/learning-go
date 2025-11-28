package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPipeline(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := []float32{5, 10, 15, 20, 25, 30, 35, 40}
	assert.Equal(t, result, collector(multiplier(generator(input))))
}
