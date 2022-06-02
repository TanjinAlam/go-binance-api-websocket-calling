package main

import (
	"context"
	"fmt"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
)


func main() {
	futures.UseTestnet = true
	var (
		apiKey = "772d7e9f365aad46411b867c5c110731a32d19fec30ee179bda21d486c9cd129"
		secretKey = "27050274c0f55263493d3a0758351e3c7e26b14d9b97dd20cd276f2b71739223"
	)
	futuresClient := binance.NewFuturesClient(apiKey, secretKey)    // USDT-M Futures
	fmt.Println(futuresClient)

	listenKey, err := futuresClient.NewStartUserStreamService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print(listenKey)
	
	// this one create issue unable to get userData
	wsDepthHandlers := func(event *binance.WsUserDataEvent) {
		fmt.Println(event)
	}
	errHandlers := func(err error) {
		fmt.Println(err)
	}
	doneCD, stopCD, err := binance.WsUserDataServe(listenKey, wsDepthHandlers, errHandlers)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopCD <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneCD

	// api calling is working fine
	// getting account balance from api call
	ress, errr := futuresClient.NewGetAccountService().Do(context.Background())
	if errr != nil {
		fmt.Println(err)
		return
	}
	val := ress.Assets
	for _, value := range val {
        fmt.Println(value)
    }
	// this websocket working is working fine
	wsDepthHandler := func(event *binance.WsMarketStatEvent) {
		fmt.Println(event)
	}
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsMarketStatServe("LTCUSDT", wsDepthHandler, errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC


}