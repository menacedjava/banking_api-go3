package main

import (
    "encoding/json"
    "net/http"
)

type Account struct {
    ID int `json:"id"`
    Balance float64 `json:"balance"`
}

var accounts = []Account{{1,1000},{2,500}}

func main() {
    http.HandleFunc("/transfer", func(w http.ResponseWriter, r *http.Request) {
        var data struct{ From, To, Amount int }
        json.NewDecoder(r.Body).Decode(&data)
        for i := range accounts {
            if accounts[i].ID == data.From { accounts[i].Balance -= float64(data.Amount) }
            if accounts[i].ID == data.To { accounts[i].Balance += float64(data.Amount) }
        }
        json.NewEncoder(w).Encode(accounts)
    })
    http.ListenAndServe(":8080", nil)
}
