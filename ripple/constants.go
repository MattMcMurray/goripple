package ripple

const (
	// RequestSuccess when a request to a rippled server succeeds, this should be the value of the
	// result.status field
	RequestSuccess = "success"
	// EngineSuccess when a request to a rippled server that interacts with the engine succeeds,
	// this should be the value of the result.engine_result field
	EngineSuccess = "tesSUCCESS"
)
