package helper

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
)

func AddString(arg ...string) float64 {
	sum := float64(0)
	for _, s := range arg {
		f, _ := strconv.ParseFloat(s, 64)

		sum += f
	}

	return sum
}

func GetValidMessage(err error, obj interface{}) string {
	getObj := reflect.TypeOf(obj)

	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") // 返回第一条 msg
			}
		}
	}

	return err.Error()
}