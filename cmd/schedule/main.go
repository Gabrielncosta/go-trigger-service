package main

import (
	"github.com/gabrielncosta/linkapi-golang/internal/orders"
	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(orders.HandlePendingOrders)
	<-s.Start()
}