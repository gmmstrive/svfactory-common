package result

import "net/http"

var (
	// Success 操作成功
	Success = Message{code: http.StatusOK, msg: "操作成功"}

	// Failure 业务异常
	Failure = Message{code: http.StatusBadRequest, msg: "业务异常"}

	// UnAUTHORIZED 请求未授权
	UnAUTHORIZED = Message{code: http.StatusUnauthorized, msg: "请求未授权"}

	// NotFound 404 没找到请求
	NotFound = Message{code: http.StatusNotFound, msg: "404 没找到请求"}

	// MsgNotReadable 消息不能读取
	MsgNotReadable = Message{code: http.StatusBadRequest, msg: "消息不能读取"}

	// MethodNotSupported 不支持当前请求方法
	MethodNotSupported = Message{code: http.StatusMethodNotAllowed, msg: "不支持当前请求方法"}

	// MediaTypeNotSupported 不支持当前媒体类型
	MediaTypeNotSupported = Message{code: http.StatusUnsupportedMediaType, msg: "不支持当前媒体类型"}

	// ReqReject 请求被拒绝
	ReqReject = Message{code: http.StatusForbidden, msg: "请求被拒绝"}

	// InternalServerError 服务器异常
	InternalServerError = Message{code: http.StatusInternalServerError, msg: "服务器异常"}

	// ParamMiss 缺少必要的请求参数
	ParamMiss = Message{code: http.StatusBadRequest, msg: "缺少必要的请求参数"}

	// ParamTypeError 请求参数类型错误
	ParamTypeError = Message{code: http.StatusBadRequest, msg: "请求参数类型错误"}

	// ParamBindError 请求参数绑定错误
	ParamBindError = Message{code: http.StatusBadRequest, msg: "请求参数绑定错误"}

	// ParamValidError 参数校验失败
	ParamValidError = Message{code: http.StatusBadRequest, msg: "参数校验失败"}
)

type Message struct {
	msg  string
	code int
}

func (r *Message) GetMessage() string {
	return r.msg
}

func (r *Message) GetCode() int {
	return r.code
}
