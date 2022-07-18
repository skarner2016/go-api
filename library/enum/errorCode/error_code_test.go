package errorCode

import (
	"fmt"
	"testing"
)

func TestErrorCode_GetMsg(t *testing.T) {
	msg := NewError().GetMessage(UnknownError)
	fmt.Println(msg)
}
