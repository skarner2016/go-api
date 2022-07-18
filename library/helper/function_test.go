package helper

import (
	"fmt"
	"testing"
)

func TestAddString(t *testing.T) {
	AddString("123", "23", "456")
}

func TestRandString(t *testing.T) {
	fmt.Println(RandString(10))
}