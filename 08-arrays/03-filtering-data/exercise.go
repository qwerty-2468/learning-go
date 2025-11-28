package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	// INSERT YOUR CODE HERE
	var result [10]string

	// If lengths differ, return array of 10 empty strings.
	if len(keys) != len(indices) {
		return result
	}

	pos := 0
	for i := range keys {
		if indices[i] > 4 {
			if pos < len(result) {
				result[pos] = keys[i]
				pos++
			} else {
				break
			}
		}
	}

	// Remaining positions (if any) stay as empty strings.
	return result
}
