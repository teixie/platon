package platon

const (
	StatusOK = 200
	MsgOK    = "success"
)

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e ApiError) Error() string {
	return e.Msg
}

type ErrorWrapper struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Response(err ApiError, args ...interface{}) interface{} {
	if len(args) == 1 {
		return ErrorWrapper{err.Code, err.Msg, args[0]}
	}
	return ErrorWrapper{err.Code, err.Msg, args}
}

func ResponseOK(args ...interface{}) interface{} {
	return Response(ApiError{StatusOK, MsgOK}, args...)
}
