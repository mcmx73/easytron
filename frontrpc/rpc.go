package frontrpc

const (
	JSON_RPC_VERSION = "2.0"

	ERROR_CODE_PARSE_ERROR         = -32700
	ERROR_MESSAGE_PARSE_ERROR      = "Parse error"
	ERROR_CODE_INVALID_REQUEST     = -32600
	ERROR_MESSAGE_INVALID_REQUEST  = "invalid request"
	ERROR_CODE_METHOD_NOT_FOUND    = -32601
	ERROR_MESSAGE_METHOD_NOT_FOUND = "method not found"
	ERROR_CODE_SERVER_ERROR        = -32000
	ERROR_MESSAGE_SERVER_ERROR     = "server error"
)

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Rpc struct {
	initComplete bool
	Id           RequestId              `json:"id"`
	Jsonrpc      string                 `json:"jsonrpc"`
	Method       string                 `json:"method"`
	Params       map[string]interface{} `json:"params,omitempty"`
}
