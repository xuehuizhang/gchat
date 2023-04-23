package enum

const (
	NormalCode       = 0
	AuthErrCode      = 10001
	SerializeErrCode = 10002
)

func GetErrorMsg(code int) string {
	switch code {
	case AuthErrCode:
		return AuthError
	default:
		return InternalError
	}
}
