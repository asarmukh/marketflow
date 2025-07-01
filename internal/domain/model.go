package domain

import "time"

type MarketData struct {
	Exchange  string
	Pair      string
	price     string
	Timestamp time.Time
}

type AggregatedData struct {
	PairName     string
	Exchange     string
	Timestamp    time.Time
	AveragePrice float64
	MaxPrice     float64
	MinPrice     float64
}

type SystemStatus struct {
	RedisConnected    bool     `json:"redis_connected"`
	PostgresConnected bool     `json:"postgres_connected"`
	ExchangeConnected []string `json:"exchange_connected"`
	Mode              string   `json:"mode"`
}
