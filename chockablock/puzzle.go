package chockablock

import (
	"errors"
	"time"

	"github.com/ianberinger/stockfighter/api"
	"github.com/kr/pretty"
)

type Config struct {
	ApiKey         string
	Account        string
	Venue          string
	Symbol         string
	SharesPerTrade int
	SharesGoal     int
	MinPrice       int
	PriceDrop      int
	PriceStart     int
}

func Run(cfg Config) error {
	i := api.NewInstance(cfg.ApiKey, cfg.Account, cfg.Venue, cfg.Symbol)

	if r := i.Heartbeat(); !r.Ok {
		return errors.New(r.Message)
	}

	if r := i.VenueHeartbeat(); !r.Ok {
		return errors.New(r.Message)
	}

	sharesAcquired := 0
	price := cfg.PriceStart
	dropCounter := 0
	for sharesAcquired < cfg.SharesGoal {
		// quote := i.Quote()

		order := i.NewOrder(price, cfg.SharesPerTrade, api.Buy, api.Limit)
		if err := i.GetErr(); err != nil {
			pretty.Println("[ERROR] NewOrder:", err)
		} else {
			pretty.Println("Created order:", order)
			//see status of order
			pretty.Println("waiting for 1 second before querying order status")
			time.Sleep(1 * time.Second)
			pretty.Println("status of order:", i.OrderStatus(order.ID))
		}

		if dropCounter == 20 && price > cfg.MinPrice {
			dropCounter = 0
			price -= cfg.PriceDrop
		}

		sharesAcquired += cfg.SharesPerTrade
		dropCounter += 1
	}

	pretty.Println("Done!")
	return nil
}
