package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type orderRequest struct {
	UUID         string  `json:"uuid"`
	Price        float64 `json:"price"`
	ContractType string  `json:"contracttype"`
}

func theBanksOpinion() orderRequest {
	usernames := []string{"User1", "User2", "User3", "User4", "User5"}
	contractTypes := []string{"Bid", "Ask"}
	return orderRequest{UUID: usernames[rand.Intn(len(usernames))], Price: rand.Float64(), ContractType: contractTypes[rand.Intn(len(contractTypes))]}
}

func main() {
	url := os.Getenv("EXCHANGE_URL")
	for {
		order := theBanksOpinion()
		jsonValue, _ := json.Marshal(order)
		time.Sleep(time.Duration(100) * time.Nanosecond)
		_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
		if err != nil {
			time.Sleep(time.Duration(5) * time.Second)
			fmt.Println(err)
		}
	}
}
