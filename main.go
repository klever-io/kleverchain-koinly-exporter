package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"
)

const (
	api           = "https://api.mainnet.klever.finance"
	txListPath    = "transaction/list"
	assetsPath    = "assets"
	kleverscanURL = "https://kleverscan.org/transaction/"
)

const csvHeader = "Date,Sent Amount,Sent Currency,Received Amount,Received Currency,Fee Amount,Fee Currency,Net Worth Amount,Net Worth Currency,Tag,Description,TxHash\n"

type KoinlyTransaction struct {
	Date             string  `json:"Date"`
	SentAmount       float64 `json:"Sent Amount"`
	SentCurrency     string  `json:"Sent Currency"`
	ReceivedAmount   float64 `json:"Received Amount"`
	ReceivedCurrency string  `json:"Received Currency"`
	FeeAmount        float64 `json:"Fee Amount"`
	FeeCurrency      string  `json:"Fee Currency"`
	NetWorthAmount   float64 `json:"Net Worth Amount"`
	NetWorthCurrency string  `json:"Net Worth Currency"`
	Tag              string  `json:"Tag"`
	Description      string  `json:"Description"`
	TxHash           string  `json:"TxHash"`
}

func (tx *KoinlyTransaction) String() string {
	return fmt.Sprintf("%s,%f,%s,%f,%s,%f,%s,%f,%s,%s,%s,%s\n",
		tx.Date, tx.SentAmount, tx.SentCurrency, tx.ReceivedAmount, tx.ReceivedCurrency,
		tx.FeeAmount, tx.FeeCurrency, tx.NetWorthAmount, tx.NetWorthCurrency,
		tx.Tag, tx.Description, tx.TxHash)
}

func main() {
	// read address from arguments
	address := os.Args[1]

	// if no address is given, request address
	if address == "" {
		fmt.Print("Enter address: ")
		_, err := fmt.Scanln(&address)
		if err != nil {
			panic(err)
		}
	}

	// validate addess
	if len(address) != 62 || address[:4] != "klv1" {
		panic("address must start with klv1 and have a length of 62")
	}

	exportList := make([]KoinlyTransaction, 0)

	// do from page 1  to end
	for page := 1; ; page++ {
		result, err := getPageFromAPI(address, 100, page)
		if err != nil {
			panic(err)
		}
		// parse result
		exportList = append(exportList, ParseToKoinlyTX(address, &result.Data)...)

		// check page
		if result.Pagination.TotalPages <= int64(page) {
			break
		}
	}

	// export to csv
	exportToCSV(address, exportList)
}

func getPageFromAPI(address string, limit int, page int) (*TxListResponse, error) {
	// build url
	url := fmt.Sprintf("%s/%s?address=%s&limit=%d&page=%d", api, txListPath, address, limit, page)
	// send request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// read response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// parse response
	var response TxListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func exportToCSV(address string, txList []KoinlyTransaction) {
	// create file with date and time and address name
	fileName := fmt.Sprintf("%s_%s.csv", address, time.Now().Format("20060102_150405"))
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// write header
	_, err = file.WriteString(csvHeader)
	if err != nil {
		panic(err)
	}
	// write tx list
	for _, tx := range txList {
		_, err = file.WriteString(tx.String())
		if err != nil {
			panic(err)
		}
	}
	// close file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}

var tokenBaseCache map[string]float64 = make(map[string]float64)

func GetTokenBase(assetId string) float64 {
	if value, exists := tokenBaseCache[assetId]; exists {
		return value
	}

	// // build url
	url := fmt.Sprintf("%s/%s/%s", api, assetsPath, assetId)
	// send request
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// read response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// parse response
	var response AssetResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	// cache value
	tokenBaseCache[assetId] = math.Pow10(response.Data.Asset.Precision)
	return tokenBaseCache[assetId]
}
