package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	// INSERT YOUR CODE HERE
	var h byte

	// Bits 7..6: packet type field "01" (bit7=0, bit6=1)
	h |= 0b0100_0000

	// Bit 3: FirstAttempt
	if isFirstAttempt {
		h |= 0b0001_0000
	}

	// Bits 2..1: QoS fixed to "Exactly once" => 10
	h |= 0b0000_1000

	// Bit 2: QoS already set; Bit1 is part of QoS.
	// Bit 2 stays 1, Bit1 stays 0 regardless of Broadcast.
	// Broadcast is defined at bit 1 in the spec, so:
	if isBroadcasted {
		h |= 0b0000_0010
	}

	// Bit 0: Secure
	if isSecure {
		h |= 0b0000_0001
	}

	return h
}
