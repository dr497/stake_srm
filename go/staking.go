package main

import (
	"log"
	. "stakeSrm/ftx"
	"sync"
	"time"
)

var COINS = []string{"SRM", "SRM_LOCKED", "MSRM", "MSRM_LOCKED", "SOL", "RAY"}

func auto_staking(client *FtxClient, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		for _, coin := range SRM_COINS {
			stake_coin(client, coin)
			time.Sleep(time.Second * 60 * 60)
		}
	}
}

func stake_coin(client *FtxClient, coin string) {
	available_balance, err := client.GetAvalaibleBalanceCoin(coin)
	if err != nil {
		log.Printf("Error getting available SRM %s", err)
	}
	if available_balance > 0 {
		_, err := client.Stakes(coin, available_balance)
		if err != nil {
			log.Printf("Error staking SRM %s", err)
		}
	}
}
