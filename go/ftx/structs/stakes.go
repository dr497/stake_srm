package structs

import (
	"time"
)

type Stakes struct {
	Coin string  `json:"coin"`
	Size float64 `json:"size"`
}

type StakesResponse struct {
	Coin      string    `json:"coin"`
	CreatedAt time.Time `json:"createdAt"`
	ID        int       `json:"id"`
	Size      float64   `json:"size"`
}
