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

func (c *Key) GetMobileVerificationCode(verificationType, areaCode, mobile int64) string {
	return fmt.Sprintf("user:verification_code:mobile:%d:%d:%d", verificationType, areaCode, mobile)
}

func (c *Key) GetEmailVerificationCode(verificationType int64, email string) string {
	return fmt.Sprintf("user:verification_code:email:%d:%s", verificationType, email)
}
