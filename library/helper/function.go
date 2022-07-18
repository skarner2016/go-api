package helper

import (
	"github.com/go-playground/validator/v10"
	"math/rand"
	"reflect"
	"strconv"
	"time"
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

func RandString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	strList := []byte(str)

	result := []byte{}

	i:=0

	r:=rand.New(rand.NewSource(time.Now().UnixNano()))

	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i += 1
	}

	return string(result)
}