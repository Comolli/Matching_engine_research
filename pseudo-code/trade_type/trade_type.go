package trade_type

import (
	"container/list"

	"github.com/micro/go-micro/util/log"
	"github.com/shopspring/decimal"
)

type orderQueue struct {
	sortBy     enum.SortDirection
	parentList *list.List
	elementMap map[string]*list.Element
}

func (q *orderQueue) init(sortBy enum.SortDirection) {
	q.sortBy = sortBy
	q.parentList = list.New()
	q.elementMap = make(map[string]*list.Element)
}

func AddOrder(Order) {

}
func GetHeadOrder() {

}
func PopHeadOrder() {

}
func RemoveOrder(order) {

}
func (q *orderQueue) getDepthPrice(depth int) (string, int) {
	if q.parentList.Len() == 0 {
		return "", 0
	}
	p := q.parentList.Front()
	i := 1
	for ; i < depth; i++ {
		t := p.Next()
		if t != nil {
			p = t
		} else {
			break
		}
	}
	o := p.Value.(*list.List).Front().Value.(*Order)
	return o.Price.String(), i
}
func dealBuyMarketTop(order *Order, book *orderBook, lastTradePrice *decimal.Decimal, depth int) {
	priceStr, _ := book.getSellDepthPrice(depth)
	if priceStr == "" {
		cancelOrder(order)
		return
	}
	limitPrice, _ := decimal.NewFromString(priceStr)
LOOP:
	headOrder := book.getHeadSellOrder()
	if headOrder != nil && limitPrice.GreaterThanOrEqual(headOrder.Price) {
		matchTrade(headOrder, order, book, lastTradePrice)
		if order.Amount.IsPositive() {
			goto LOOP
		}
	} else {
		cancelOrder(order)
	}
}
func dealBuyMarketOpponent(order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
	priceStr, _ := book.getSellDepthPrice(1)
	if priceStr == "" {
		cancelOrder(order)
		return
	}
	limitPrice, _ := decimal.NewFromString(priceStr)
LOOP:
	headOrder := book.getHeadSellOrder()
	if headOrder != nil && limitPrice.GreaterThanOrEqual(headOrder.Price) {
		matchTrade(headOrder, order, book, lastTradePrice)
		if order.Amount.IsPositive() {
			goto LOOP
		}
	} else {
		order.Price = limitPrice
		order.Type = enum.TypeLimit
		book.addBuyOrder(order)
		cache.UpdateOrder(order.ToMap())
		log.Info("engine %s, a order has added to the orderbook: %s", order.Symbol, order.ToJson())
	}
}
