package model

import "time"

type ResBodyReportTransaction struct {
	ID         string
	Amount     int64
	UserId     string
	MerchantId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
