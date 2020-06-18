package main

import (
	"fmt"
	"github.com/NataliaOleinik/Transactions/pkg/card"
	_ "go/types"
	"time"
)

func main() {
	master:= &card.Card{
		Id: 13,
		Issuer: "MasterCard",
		Balance: 10_000_00,
		Currency: "RUB",
		Number: "5947_5473_9502_1625",
		Transaction: []*card.Transaction{
			&card.Transaction{
				Id: 1,
				AmountTransaction: 73555,
				Date: time.Now().Unix(),
				CodeMcc: []*card.CodeMcc{
					&card.CodeMcc{
						Id: "5411",
					},
				},
				Status: "Операция в обработке",
			},
			&card.Transaction{
				Id: 2,
				AmountTransaction: 120391,
				Date:  time.Now().Unix(),
				CodeMcc: []*card.CodeMcc{
					&card.CodeMcc{
						Id: "5813",
					},
				},
				Status:  "Операция в обработке",
			},

		},
	}
	fmt.Println(master)
}
