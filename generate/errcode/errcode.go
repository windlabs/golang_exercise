//go:generate stringer -type ErrCode -linecomment -output code_string.go
package errcode

type ErrCode int

// 定义错误码
const (
	ERR_CODE_OK             ErrCode = 0 // OK
	ERR_CODE_INVALID_PARAMS ErrCode = 1 // 无效参数
	ERR_CODE_TIMEOUT        ErrCode = 2 // 超时
	// ...
)

//// 定义错误码与描述信息的映射
//var mapErrDesc = map[ErrCode]string{
//	ERR_CODE_OK:             "OK",
//	ERR_CODE_INVALID_PARAMS: "无效参数",
//	ERR_CODE_TIMEOUT:        "超时",
//	// ...
//}
//
//// 根据错误码返回描述信息
//func getDescription(errCode ErrCode) string {
//	if desc, exist := mapErrDesc[errCode]; exist {
//		return desc
//	}
//
//	return fmt.Sprintf("error code: %d", errCode)
//}

//func (e ErrCode) String() string {
//	return getDescription(e)
//}
