package model

import "time"

type Merchant struct {
	ID        string
	Name      string
	Phone     string
	Address   string
	Balance   int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
