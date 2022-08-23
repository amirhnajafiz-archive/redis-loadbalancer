package response

import "time"

type PairResponse struct {
	Key   string    `json:"key"`
	Value string    `json:"value"`
	Date  time.Time `json:"date"`
}
