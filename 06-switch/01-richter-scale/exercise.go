package richterscale

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// describeEarthquake returns the "description" of a given magnitude value on the Richter scale.
func describeEarthquake(magnitude float32) string {
	// INSERT YOUR CODE HERE
	switch {
	case magnitude < 2.0:
		return "micro"
	case magnitude < 3.0:
		return "very minor"
	case magnitude < 4.0:
		return "minor"
	case magnitude < 5.0:
		return "light"
	case magnitude < 6.0:
		return "moderate"
	case magnitude < 7.0:
		return "strong"
	case magnitude < 8.0:
		return "major"
	case magnitude < 10.0:
		return "great"
	default:
		return "massive"
	}
}
