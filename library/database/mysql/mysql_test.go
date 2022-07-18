package mysql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-api/library/config"
	"testing"
)

func TestNewMysql(t *testing.T) {
	config.InitConfig()

	db, err := NewMysql(InstantDefault)
	assert.Nil(t, err)
	db2, err := NewMysql(InstantDefault)

	fmt.Println(fmt.Sprintf("%p", db))
	fmt.Println(fmt.Sprintf("%p", db2))

	//user := new(User)
	//db.First(&user)
	//fmt.Println(user)
}