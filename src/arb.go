package src

import (
	"fmt"
)

func CheckArbitrage() error {
	binanceBTCPrice, err := GetBinanceBTCToUSDTPrice()
	if err != nil {
		return fmt.Errorf("error fetching Binance BTC price: %v", err)
	}

	arbAmount := defaultArbAmt
	fmt.Println("Binance BTC Price:", binanceBTCPrice)

	osmosisBTCPrice, route, err := GetOsmosisBTCToUSDCPriceAndRoute(arbAmount)
	if err != nil {
		return fmt.Errorf("error fetching Osmosis BTC price: %v", err)
	}

	fmt.Println("Osmosis BTC Price:", osmosisBTCPrice)

	if binanceBTCPrice < osmosisBTCPrice*riskFactor {
		fmt.Println("Arbitrage Opportunity: Buy BTC on Binance, Sell BTC on Osmosis")

		_, route, err := GetOsmosisUSDCToBTCPriceAndRoute(arbAmount)

		if err != nil {
			return err
		}

		err = SellOsmosisBTC(seedConfig, route, binanceBTCPrice)
		if err != nil {
			return err
		}

		_, _, err = BuyBinanceBTC(arbAmount)
		if err != nil {
			return err
		}

	} else if binanceBTCPrice*riskFactor > osmosisBTCPrice {
		fmt.Println("Arbitrage Opportunity: Sell BTC on Binance, Buy BTC on Osmosis")

		err = BuyOsmosisBTC(seedConfig, route, binanceBTCPrice)
		if err != nil {
			return err
		}

		_, _, err = SellBinanceBTC(arbAmount)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("no arb opportunity")
	}

	return nil
}

// b, u, err := GetBinanceBTCUSDTBalance()
// if err != nil {
// 	return fmt.Errorf("error fetching Osmosis BTC price: %v", err)
// }

// fmt.Printf("btc, usdc balance: %f, %f\n", b, u)

// boughtAmt, _, err := BuyBinanceBTC(defaultArbAmt)
// if err != nil {
// 	return err
// }
// fmt.Printf("bought %f\n", boughtAmt)

// b, u, err = GetBinanceBTCUSDTBalance()
// if err != nil {
// 	return fmt.Errorf("error fetching Osmosis BTC price: %v", err)
// }

// fmt.Printf("btc, usdc balance: %f, %f\n", b, u)

// sellAmt, _, err := SellBinanceBTC(defaultArbAmt)
// if err != nil {
// 	return err
// }
// fmt.Printf("sold %f\n", sellAmt)
