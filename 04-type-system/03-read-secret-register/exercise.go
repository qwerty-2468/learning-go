package readsecretregister

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// parseChannelControlRegister constructs 4 octets (8-bit long uint) based on the parameter register.
func parseChannelControlRegister(charCtrl uint32) (uint8, uint8, uint8, uint8) {
	// INSERT YOUR CODE HERE
	rxPcode := uint8((charCtrl >> 24) & 0xFF)
	txPcode := uint8((charCtrl >> 16) & 0xFF)
	rxChan := uint8((charCtrl >> 8) & 0xFF)
	txChan := uint8(charCtrl & 0xFF)

	// Return in order: TX_CHAN, RX_CHAN, RX_PCODE, TX_PCODE
	return txChan, rxChan, rxPcode, txPcode
}
