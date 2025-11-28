package calculator

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE

import (
	"math"
	"strconv"
)

func gmean(x, y float64) int {
    gm := math.Sqrt(x * y)
    return int(math.Round(gm))
}

func gmeanString(x, y string) (int, error) {
    xf, err := strconv.ParseFloat(x, 64)
    if err != nil {
        return 0, err
    }

    yf, err := strconv.ParseFloat(y, 64)
    if err != nil {
        return 0, err
    }

    return gmean(xf, yf), nil
}