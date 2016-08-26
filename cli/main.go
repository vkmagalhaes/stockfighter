package main

import (
	"errors"
	"flag"
	"log"

	"github.com/vkmagalhaes/stockfighter/chockablock"
)

func main() {
	var (
		// program flags
		puzzle  = flag.String("puzzle", "", "cmd puzzles: chockablock")
		apiKey  = flag.String("api-key", "", "Stockfighter api key")
		account = flag.String("account", "", "puzzle's account")
		venue   = flag.String("venue", "", "puzzle's venue")
		symbol  = flag.String("symbol", "", "puzzle's symbol")

		sharesPerTrade = flag.Int("shares.per-trade", 100, "number of shares to be executed per trade")
		sharesGoal     = flag.Int("shares.goal", 100000, "number of shares to be executed in total")
		priceMin       = flag.Int("price.min", 100, "min price to negotiate the asset")
		priceDrop      = flag.Int("price.drop", 50, "forced price decrementer")
		priceStart     = flag.Int("price.start", 5000, "forced price decrementer")
	)

	flag.Parse()

	var err error
	switch *puzzle {
	case "chockablock":
		err = chockablock.Run(chockablock.Config{
			ApiKey:         *apiKey,
			Account:        *account,
			Venue:          *venue,
			Symbol:         *symbol,
			SharesPerTrade: *sharesPerTrade,
			SharesGoal:     *sharesGoal,
			MinPrice:       *priceMin,
			PriceDrop:      *priceDrop,
			PriceStart:     *priceStart,
		})
	default:
		err = errors.New("puzzle not implemented")
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Done!")
}
