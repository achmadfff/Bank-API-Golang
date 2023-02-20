package model

type ReqBodyTransaction struct {
	Amount     int64  `binding:"required"`
	MerchantId string `binding:"required"`
}
