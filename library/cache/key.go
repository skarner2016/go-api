package cache

import "fmt"

type Key struct {

}

func NewKey() *Key {
	return &Key{}
}

func (c *Key) GetSymbolPrice() string {
	return "symbol:price:binance"
}

func (c *Key) GetSymbolInfo() string {
	return "symbol:binance"
}

func (c *Key) GetMobileVerificationCode(areaCode, mobile int64) string {
	return fmt.Sprintf("user:verification_code:mobile:%d:%d", areaCode, mobile)
}

func (c *Key) GetEmailVerificationCode(email string) string {
	return fmt.Sprintf("user:verification_code:email:%s", email)
}
