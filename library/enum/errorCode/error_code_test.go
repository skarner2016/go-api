package errorCode

import (
	"fmt"
	"testing"
)

func TestErrorCode_GetMsg(t *testing.T) {
	msg := NewError().GetMsg(UnknownError)
	fmt.Println(msg)
}
