package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorReqponse struct {
	HttpSc int
	Error Err
}

// 初始化ErrorReqponse
var (
	// body无法解析
	ErrorRequestsBodyParseFaild = ErrorReqponse{400, Err{"Request body is not correct", "001"}}
	// 用户验证不通过
	ErrorNotAuthUser = ErrorReqponse{401, Err{"User authentication failed.", "002"}}
)