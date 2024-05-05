package pkg

import (
	"sync"
	"time"
)

type RequestInfo struct {
	ID   int
	Time time.Time
}

type responseRequest struct {
	ID     int         `json:"id"`
	Result interface{} `json:"result"`
}

type generateRequest struct {
	Type   string `json:"type"`
	Length int    `json:"length"`
}

var requestStore = make(map[string]RequestInfo)
var store = make(map[int]interface{})
var nextID = 0
var mu sync.Mutex
