package main

import (
	"time"
)

type AssetResponse struct {
	Data struct {
		Asset struct {
			AssetType string `json:"assetType"`
			AssetID   string `json:"assetId"`
			Name      string `json:"name"`
			Ticker    string `json:"ticker"`
			Precision int    `json:"precision"`
		} `json:"asset"`
	} `json:"data"`
}

type TXContract struct {
	Type       TXContract_ContractType `json:"type"`
	TypeString string                  `json:"typeString"`
	Parameter  interface{}             `json:"parameter,omitempty"`
}

type Transaction struct {
	Hash         string                   `json:"hash"`
	BlockNum     uint64                   `json:"blockNum,omitempty"`
	Sender       string                   `json:"sender"`
	Nonce        uint64                   `json:"nonce"`
	PermissionID int32                    `json:"permissionID,omitempty"`
	Data         []string                 `json:"data,omitempty"`
	Timestamp    time.Duration            `json:"timestamp,omitempty"`
	KAppFee      int64                    `json:"kAppFee"`
	KDAFee       *KDAFee                  `json:"kdaFee,omitempty"`
	BandwidthFee int64                    `json:"bandwidthFee"`
	Status       string                   `json:"status"`
	ResultCode   string                   `json:"resultCode,omitempty"`
	Version      uint32                   `json:"version,omitempty"`
	ChainID      string                   `json:"chainID,omitempty"`
	Signature    []string                 `json:"signature,omitempty"`
	SearchOrder  uint32                   `json:"searchOrder"`
	Receipts     []map[string]interface{} `json:"receipts"`
	Contracts    []*TXContract            `json:"contract"`
}

type TxListResponseResponseData struct {
	Transactions []*Transaction `json:"transactions"`
}

type Pagination struct {
	Self         int64 `json:"self"`
	Next         int64 `json:"next"`
	Previous     int64 `json:"previous"`
	PerPage      int64 `json:"perPage"`
	TotalPages   int64 `json:"totalPages"`
	TotalRecords int64 `json:"totalRecords"`
}

// TxListResponse structure
type TxListResponse struct {
	Data       TxListResponseResponseData `json:"data"`
	Pagination *Pagination                `json:"pagination,omitempty"`
	Error      string                     `json:"error"`
	Code       string                     `json:"code"`
}

type KDAFee struct {
	KDA    string `json:"kda"`
	Amount int64  `json:"amount"`
}
