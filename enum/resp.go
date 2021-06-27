package enum

type RespStatus int

//	响应枚举
const (
	SUCCESS RespStatus = 200
	ERROR   RespStatus = -200
)

//	响应枚举方法
func (respStatus RespStatus) String() string {
	switch respStatus {
	case SUCCESS:
		return "success"
	case ERROR:
		return "error"
	default:
		return "unknown"
	}
}
