package ftx

import (
	"log"
	"stakeSrm/ftx/structs"
)

type Balances structs.Balances

func (client *FtxClient) GetBalances() (Balances, error) {
	var balances Balances
	resp, err := client._get("wallet/balances", []byte(""))
	if err != nil {
		log.Printf("Error GetBalances %s", err)
		return balances, err
	}
	err = ProcessResponse(resp, &balances)
	return balances, err
}

func (client *FtxClient) GetAvalaibleBalanceCoin(coin string) (float64, error) {
	var balance float64
	resp, err := client.GetBalances()
	if err != nil {
		log.Printf("Error GetBalance %s", err)
		return balance, err
	}
	for _, v := range resp.Result {
		if v.Coin == coin {
			return v.Free, nil
		}
	}
	return balance, err
}
