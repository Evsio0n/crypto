package cryptos

/*
	THIS PACKAGE HAS ABSOLUTELY NO WARRANTY.
*/

import (
	"github.com/go-resty/resty/v2"
	"time"
)

var usdtTrc20Address = "TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t"
var tronGridUri = "https://apiasia.trongrid.io:5566/api/token_trc20/transfers"
var client *resty.Client

var tronGridUriParams = map[string]string{
	"direction":       "in",
	"count":           "8",
	"tokens":          usdtTrc20Address,
	"relatedAddress":  "",
	"start_timestamp": "",
}

type TrongridResponse struct {
	Total        int `json:"total"`
	RangeTotal   int `json:"rangeTotal"`
	ContractInfo struct {
	} `json:"contractInfo"`
	TokenTransfers []struct {
		TransactionId   string `json:"transaction_id"`
		BlockTs         int64  `json:"block_ts"`
		FromAddress     string `json:"from_address"`
		ToAddress       string `json:"to_address"`
		Block           int    `json:"block"`
		ContractAddress string `json:"contract_address"`
		Quant           string `json:"quant"`
		Confirmed       bool   `json:"confirmed"`
		ContractRet     string `json:"contractRet"`
		FinalResult     string `json:"finalResult"`
		Revert          bool   `json:"revert"`
		TokenInfo       struct {
			TokenId      string `json:"tokenId"`
			TokenAbbr    string `json:"tokenAbbr"`
			TokenName    string `json:"tokenName"`
			TokenDecimal int    `json:"tokenDecimal"`
			TokenCanShow int    `json:"tokenCanShow"`
			TokenType    string `json:"tokenType"`
			TokenLogo    string `json:"tokenLogo"`
			TokenLevel   string `json:"tokenLevel"`
			IssuerAddr   string `json:"issuerAddr"`
			Vip          bool   `json:"vip"`
		} `json:"tokenInfo"`
		FromAddressIsContract bool `json:"fromAddressIsContract"`
		ToAddressIsContract   bool `json:"toAddressIsContract"`
	} `json:"token_transfers"`
}

// GetUSDTTrc20Transfer is a function to get USDT Trc 20 transfer from trongrid api .
//
// @param startTimeStamp is the start time of the query.
//
// @param amount is the amount of usdt.
//
// @param address is the address of the usdt in trc20 format. NOT A CONTRACT ADDRESS OR HEX FORMAT.
//
// This should be verified with timestamp. if user had sent a same amount of USDT before, it will be ignored.
func GetUSDTTrc20Transfer(USDTAddress string, startTimeStamp int64, amount int64) bool {
	tronGridRequest := tronGridUriParams
	tronGridRequest["tokens"] = USDTAddress
	tronGridRequest["start_timestamp"] = string(startTimeStamp * 1000)

	client.SetPathParams(tronGridRequest)

	var response TrongridResponse
	//Time out set to 10
	client.SetTimeout(time.Second * 10)
	client.R().SetResult(&response).Get(tronGridUri)
	for _, Transfer := range response.TokenTransfers {
		if Transfer.Quant == string(amount*1000000) {
			return true
		}
	}
	return false
}
