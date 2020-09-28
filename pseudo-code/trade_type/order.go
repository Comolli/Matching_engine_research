package trade_type

import (
	"encoding/json"
	"time"

	"github.com/shopspring/decimal"
)

const (
	create = iota
	cancel
	buy
	sell
	limit
	market
)

type Order struct {
	action    uint8 `json:"action"`
	side      uint8
	orderId   uint64
	typeTrade uint
	amount    uint64
	price     *decimal.Decimal
	symbol    string
	timeStamp time.Time
}

func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

func (order *Order) ToJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}

type Trade struct {
	makerId   uint64
	takerId   uint64
	amount    uint64
	takerSide int8
	price     *decimal.Decimal
	timeStamp time.Time
}
