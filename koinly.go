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
	kList := make([]KoinlyTransaction, 0)

	for _, tx := range txList.Transactions {
		date := time.Unix(int64(tx.Timestamp), 0).UTC().Format("2006-01-02 15:04:05")
		fee := float64(tx.BandwidthFee+tx.KAppFee) / GetTokenBase("KLV")
		// TODO: check for token fee != KLV
		feeCurrency := "KLV"

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

			// only have fee if sender == tx.Sender
			if sender != tx.Sender {
				fee = 0
				feeCurrency = ""
			}

			if kt := decodeTransaction(tx.Hash, tx.Sender, sender, c, tx.Receipts); kt != nil {
				kt.Date = date

				if kt.SentAmount == 0 && kt.ReceivedAmount == 0 {
					// consider using TX fee as withdraw and tag as costs if no amount is set
					kt.SentAmount = fee
					kt.SentCurrency = "KLV"
					kt.Tag = "cost"
					description = fmt.Sprintf("Transaction Fee ContractType: %s - Index: %d", c.TypeString, idx)
				}

				kt.FeeAmount = fee
				kt.FeeCurrency = feeCurrency
				kt.Description = description
				kt.TxHash = fmt.Sprintf("%s%s", kleverscanURL, tx.Hash)

				kList = append(kList, *kt)
			}
		}
	}

	return kList
}

func decodeTransaction(hash, txSender, sender string, c *TXContract, r []map[string]interface{}) *KoinlyTransaction {
	switch c.Type {
	case 0: // transfer
		return decodeTransfer(txSender, sender, c)
	case 5: // unfreeze
		// register claim during unfreeze if any
		return decodeClaim(txSender, sender, c, r, hash)
	case 9: // claim
		return decodeClaim(txSender, sender, c, r, hash)
	}
	return nil
}

func decodeTransfer(txSender, sender string, c *TXContract) *KoinlyTransaction {
	// check if address is sender or receiver
	if txSender == sender {
		amount := float64(0)
		assetId := ""
		worthAmount := float64(0)
		worthCurrency := ""

		// collect amount and assetId

		if r_assetId, ok := c.Parameter.(map[string]interface{})["assetId"].(string); ok {
			assetId = r_assetId
		} else {
			assetId = "KLV"
		}

		amount = float64(c.Parameter.(map[string]interface{})["amount"].(float64)) / GetTokenBase(assetId)
		assetId, worthAmount, worthCurrency = swapAssetId(assetId, amount)

		return &KoinlyTransaction{
			SentAmount:       amount,
			SentCurrency:     assetId,
			ReceivedAmount:   0,
			ReceivedCurrency: "",
			NetWorthAmount:   worthAmount,
			NetWorthCurrency: worthCurrency,
		}
	}

	if c.Parameter.(map[string]interface{})["toAddress"].(string) == sender {
		// get token and decimate amount
		var assetId string
		if r_assetId, ok := c.Parameter.(map[string]interface{})["assetId"].(string); ok {
			assetId = r_assetId
		} else {
			assetId = "KLV"
		}

		amount := float64(c.Parameter.(map[string]interface{})["amount"].(float64)) / GetTokenBase(assetId)
		assetId, worthAmount, worthCurrency := swapAssetId(assetId, amount)

		return &KoinlyTransaction{
			SentAmount:       0,
			SentCurrency:     "",
			ReceivedAmount:   amount,
			ReceivedCurrency: assetId,
			NetWorthAmount:   worthAmount,
			NetWorthCurrency: worthCurrency,
		}
	}

	return nil
}

func decodeClaim(txSender, sender string, c *TXContract, receipts []map[string]interface{}, hash string) *KoinlyTransaction {
	// TODO: count for multiple tokens claimed
	amount := float64(0)
	tokenReceived := ""
	for _, r := range receipts {
		if r["type"].(float64) != 17 {
			continue
		}

		if cts, ok := r["claimType"]; (ok && cts != nil &&
			(cts.(float64) == 0 ||
				cts.(float64) == 1)) ||
			// OLD receipt format
			!ok {
			tokenReceived = "KLV"
			// OLD receipt format no tokenReceived
			if t, ok := r["assetIdReceived"]; ok && t != nil {
				tokenReceived = t.(string)
			}
			// get token and decimate amount
			v := r["amount"].(float64) / GetTokenBase(tokenReceived)
			amount += v
		}
	}

	tag := "mining"
	if amount == 0 {
		tag = "costs"
	}

	return &KoinlyTransaction{
		SentAmount:       0,
		SentCurrency:     "",
		ReceivedAmount:   amount,
		ReceivedCurrency: tokenReceived,
		NetWorthAmount:   amount,
		NetWorthCurrency: tokenReceived,
		Tag:              tag,
	}
}
