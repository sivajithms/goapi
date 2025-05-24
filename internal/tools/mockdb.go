package tools 

import (
	"time"
)

type MockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"sivajith" : {
		AuthToken : "123456",
		UserName: "sivajith",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"sivajith" : {
		Coins: 100,
		UserName: "sivajith",
	},
}

func (db *MockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *MockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (db *MockDB) SetupDatabase() error {
	return nil
}