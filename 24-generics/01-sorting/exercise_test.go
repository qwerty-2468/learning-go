package sorting

import (
	
	"math"
	
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSorting(t *testing.T) {
	

	
	arrTypeOne := []int64{3,-1,22,-11,8080,5432,123,53,179,-66}
	arrTypeTwo := []float64{1.1,-0.1,66,-22,0,math.Pi,6.2834, math.E}
	assert.Equal(t, []int64{-66,-11,-1,3,22,53,123,179,5432,8080},sortSlice(arrTypeOne))
	assert.Equal(t, []float64{-22,-0.1,0,1.1, math.E, math.Pi,6.2834,66}, sortSlice(arrTypeTwo))
	

	

	

	
}
