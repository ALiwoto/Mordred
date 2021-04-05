package operations

// the result of an operation.
type RESULT uint8


const (
	// CANCELED result.
	// it means operation was canceled!
	CANCELED RESULT = 2
	// success result.
	// it means operation was successful!
	SUCCESS RESULT = 1
	// failed result.
	// it means operation was NOT successful!
	FAILED RESULT = 0
)