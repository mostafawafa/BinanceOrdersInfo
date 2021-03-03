package main

import (
	"strings"
)

func getOrdersInfo(currency string) *orderInfo {

	var orders []order

	pairs := getPairs(currency)

	for _, pair := range pairs {
		orders = append(orders, sendRequest("allOrders", pair)...)
	}

	paid, quantity := 0.0, 0.0

	for _, order := range orders {
		if order.Status == "FILLED" {
			if order.Side == "BUY" {
				paid += order.Price * order.Quantity
				quantity += order.Quantity
			}
			if order.Side == "SELL" {
				paid -= order.Price * order.Quantity
				quantity -= order.Quantity
			}
		}
	}

	return createOrderInfo(currency, quantity, paid)

}

func getPairs(currency string) []string {
	return []string{
		strings.ToUpper(currency) + "USDT",
		strings.ToUpper(currency) + "BUSD",
	}
}

func createOrderInfo(currency string, quantity float64, paid float64) *orderInfo {
	orderInfo := &orderInfo{}
	orderInfo.Currency = currency
	orderInfo.Quantity = quantity
	orderInfo.AveragePrice = paid / quantity
	return orderInfo
}
