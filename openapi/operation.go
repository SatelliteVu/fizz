package openapi

// OperationInfo represents the informations of an operation
// that will be used when generating the OpenAPI specification.
type OperationInfo struct {
	ID                string
	StatusCode        int
	StatusDescription string
	Headers           []*ResponseHeader
	Summary           string
	Description       string
	Deprecated        bool
	InputModel        interface{}
	Responses         []*OperationResponse
	Security          []*SecurityRequirement
	XCodeSamples      []*XCodeSample
	XInternal         bool
	ExtraParams       []*ExtraParam // Additional parameters not derived from InputModel
}

// ExtraParam represents an additional parameter to include in the operation.
// This is useful for path parameters that come from parent routes (e.g., /v1/:contract_id/...)
// where the contract_id is not part of the handler's input model.
type ExtraParam struct {
	Name        string // Parameter name (e.g., "contract_id")
	In          string // Parameter location: "path", "query", "header", "cookie"
	Description string // Optional description
	Required    bool   // Whether the parameter is required (always true for path params)
}

// ResponseHeader represents a single header that
// may be returned with an operation response.
type ResponseHeader struct {
	Name        string
	Description string
	Model       interface{}
}

// OperationResponse represents a single response of an
// API operation.
type OperationResponse struct {
	// The response code can be "default"
	// according to OAS3 spec.
	Code        string
	Description string
	Model       interface{}
	Headers     []*ResponseHeader
	Example     interface{}
	Examples    map[string]interface{}
}
