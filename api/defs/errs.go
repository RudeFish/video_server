package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error Err
}

// 初始化ErrorReqponse
var (
	// body无法解析
	ErrorRequestsBodyParseFaild = ErrorResponse{400, Err{"Request body is not correct", "001"}}
	// 用户验证不通过
	ErrorNotAuthUser = ErrorResponse{401, Err{"User authentication failed.", "002"}}
	// DBErr
	ErrorDBError = ErrorResponse{500, Err{"DB ops failed", "003"}}
	// 内部错误
	ErrorInternalFaults = ErrorResponse{500, Err{"Internal service error", "004"}}
)