package ftx

import (
	"encoding/json"
	"log"
	"stakeSrm/ftx/structs"
)

type StakesResponse structs.StakesResponse
type Stakes structs.Stakes

func (client *FtxClient) Stakes(coin string, size float64) (StakesResponse, error) {
	var stakesResponse StakesResponse
	stakes, err := json.Marshal(Stakes{Coin: coin, Size: size})
	if err != nil {
		log.Printf("Error Stakes %s", err)
		return stakesResponse, err
	}
	resp, err := client._post("srm_stakes/stakes", stakes)
	if err != nil {
		log.Printf("Error Stakes %s", err)
		return stakesResponse, err
	}
	err = ProcessResponse(resp, &stakesResponse)
	return stakesResponse, err
}
