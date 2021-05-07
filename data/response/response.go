package response

type DataResponse struct {
	Status int `json:"status"`;
	Message string `json:"message"`;
	Data string `json:"data"`;
	Error ErrorStruct `json:"error"`;
}

type ValidateError struct {
	FailedField string
	Tag         string
	Value       string
}

type ErrorData struct {
	Field string;
	Tag string;
}

type ErrorStruct []ErrorData;