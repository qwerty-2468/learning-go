package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	// INSERT YOUR CODE HERE
	var h byte

	h |= 0b0100_0000

	if isFirstAttempt {
		h |= 0b0001_0000
	}


	h |= 0b0000_1000


	if isBroadcasted {
		h |= 0b0000_0010
	}

	if isSecure {
		h |= 0b0000_0001
	}

	return h
}
