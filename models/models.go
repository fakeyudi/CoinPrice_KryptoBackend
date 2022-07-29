package models

import(
	"github.com/dgrijalva/jwt-go"   //JWT Package
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserID 	int64
	Name   string
	Email  string
	*jwt.StandardClaims
}

//Exception struct declaration
type Exception struct {
	Message string `json:"message"`
}

type Trade struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	LastPrice          string `json:"c"`
	CloseQty           string `json:"Q"`
	BidPrice           string `json:"b"`
	BidQty             string `json:"B"`
	AskPrice           string `json:"a"`
	AskQty             string `json:"A"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	Count              int64  `json:"n"`
}

type Coin struct  {
	Symbol		string `json:"symbol"`
	Price		string `json:"price"`
}

type Alert struct {
	ID 		int64 		`json:"alertid"`
	UserID 	int64 		`json:"userid"`
	Symbol	string		`json:"symbol"`
	Price	string		`json:"price"`	
}