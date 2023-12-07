package main

import (
	"fmt"
	"time"
)

func swapAssetId(assetId string, amount float64) (string, float64, string) {

	switch assetId {
	case "SAT3-1EKU":
		return "SATS", amount / 100000000, "BTC"
	}

	return assetId, 0, ""
}

func ParseToKoinlyTX(sender string, txList *TxListResponseResponseData) []KoinlyTransaction {
	// TODO: check for token fee != KLV

	kList := make([]KoinlyTransaction, 0)

	for _, tx := range txList.Transactions {
		date := time.Unix(int64(tx.Timestamp), 0).UTC().Format("2006-01-02 15:04:05")
		fee := float64(tx.BandwidthFee+tx.KAppFee) / GetTokenBase("KLV")

		// if tx is from sender and it is a multsign tx, add a entry for fee
		if tx.Sender == sender && len(tx.Contracts) > 1 {
			kList = append(kList, KoinlyTransaction{
				Date:             date,
				SentAmount:       fee,
				SentCurrency:     "KLV",
				ReceivedAmount:   0,
				ReceivedCurrency: "",
				FeeAmount:        0,
				FeeCurrency:      "",
				Tag:              "cost",
				Description:      "transaction fee for multi contract transaction",
				TxHash:           fmt.Sprintf("%s%s", kleverscanURL, tx.Hash),
			})

			fee = 0
		}

		for idx, c := range tx.Contracts {
			description := fmt.Sprintf("ContractType: %s - Index: %d", c.TypeString, idx)

			// check if address is sender or receiver
			if tx.Sender == sender {
				amount := float64(0)
				assetId := ""
				worthAmount := float64(0)
				worthCurrency := ""

				// if transfer collect amount and assetId
				if c.Type == 0 {
					assetId = c.Parameter.(map[string]interface{})["assetId"].(string)
					amount = float64(c.Parameter.(map[string]interface{})["amount"].(float64)) / GetTokenBase(assetId)
					assetId, worthAmount, worthCurrency = swapAssetId(assetId, amount)
				}

				kList = append(kList, KoinlyTransaction{
					Date:             date,
					SentAmount:       amount,
					SentCurrency:     assetId,
					ReceivedAmount:   0,
					ReceivedCurrency: "",
					NetWorthAmount:   worthAmount,
					NetWorthCurrency: worthCurrency,
					FeeAmount:        fee,
					FeeCurrency:      "KLV",
					Description:      description,
					TxHash:           fmt.Sprintf("%s%s", kleverscanURL, tx.Hash),
				})
				continue
			}

			// check if transfer transaction and address is receiver
			if c.Type == 0 &&
				c.Parameter.(map[string]interface{})["toAddress"].(string) == sender {
				// get token and decimate amount
				assetId := c.Parameter.(map[string]interface{})["assetId"].(string)
				amount := float64(c.Parameter.(map[string]interface{})["amount"].(float64)) / GetTokenBase(assetId)
				assetId, worthAmount, worthCurrency := swapAssetId(assetId, amount)

				kList = append(kList, KoinlyTransaction{
					Date:             date,
					SentAmount:       0,
					SentCurrency:     "",
					ReceivedAmount:   amount,
					ReceivedCurrency: assetId,
					NetWorthAmount:   worthAmount,
					NetWorthCurrency: worthCurrency,
					Description:      description,
					TxHash:           fmt.Sprintf("%s%s", kleverscanURL, tx.Hash),
				})
			}

		}
	}

	return kList
}
