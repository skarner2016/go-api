package errorCode

import "fmt"

type Error struct {
}

type ErrorCode int64

const (
	UnknownError ErrorCode = 500
	NetWorkError ErrorCode = 500

	SendMessageFail   ErrorCode = 10001
	RegisterTypeError           = 10002
	RegisterCodeError           = 10003
	RegisterFail                = 10004
	LoginFail                   = 10005
	LoginCodeError = 10006
	DelRegisterCodeError = 10007
)

func NewError() *Error {
	return &Error{}
}

func (e *Error) GetMessage(code ErrorCode) string {
	return fmt.Sprint("error code:%d", code)
}
