package errorCode

import "fmt"

type Error struct {

}

type ErrorCode int64

const (
	UnknownError ErrorCode = 500
	NetWorkError ErrorCode = 500
)

func NewError() *Error {
	return &Error{}
}

func (e *Error) GetMsg(code ErrorCode) string {
	return fmt.Sprint("error code:%d", code)
}
