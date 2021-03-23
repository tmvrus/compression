package main

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/valyala/gozstd"
)

func main() {
	balance := AffiliateBalance{
		ClientId:    1000,
		AffiliateId: 2000,
		Items: []*Balance{
			{
				Type:     BalanceType_BALANCE_TYPE_HOLD,
				Value:    10,
				Currency: Currency_CURRENCY_USD,
			},
			{
				Type:     BalanceType_BALANCE_TYPE_CONFIRMED,
				Value:    20,
				Currency: Currency_CURRENCY_USD,
			},
		},
	}

	data, err := proto.Marshal(&balance)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Protobuf: %d\n", len(data))

	data, err = json.Marshal(&balance)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Jsoned protobuf: %d\n", len(data))

	fmt.Printf("rawJSON: %d\n", len(rawJson))

	zJson := gozstd.Compress(nil, rawJson)

	fmt.Printf("Zstandart: %d\n", len(zJson))
}

var rawJson = []byte(`{"client_id":1000,"affiliate_id":2000,"balance":[{"currency":"USD","type":"hold","value":10},{"currency":"USD","type":"confirmed","value":20}]}`)
