package calculator

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type OperationType int

const (
	Add OperationType = iota
	Subtract
	Multiply
)

// String returns the textual name of the operation.
func (op OperationType) String() string {
	switch op {
	case Add:
		return "Add"
	case Subtract:
		return "Subtract"
	case Multiply:
		return "Multiply"
	default:
		return "Unknown"
	}
}

// Calculate applies the given operation to a and b.
func Calculate(op OperationType, a, b float64) float64 {
	switch op {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	default:
		return 0
	}
}