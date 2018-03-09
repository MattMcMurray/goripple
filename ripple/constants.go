package ripple

const (
	// RequestSuccess when a request to a rippled server succeeds, this should be the value of the
	// result.status field
	RequestSuccess = "success"
	// EngineSuccess when a request to a rippled server that interacts with the engine succeeds,
	// this should be the value of the result.engine_result field
	EngineSuccess = "tesSUCCESS"

	// NoDestination is the engine message returned when a wallet does not have sufficient XRP to exist.
	NoDestination = "Destination does not exist. Send XRP to create it."

	// ServerOverloaded the body of a request returned when rippled instance cannot handle request
	ServerOverloaded = "Server is overloaded"

	// XactionExists is thrown by the server when a transaction hash is submitted more than once
	XactionExists = "The exact transaction was already in this ledger."
)
