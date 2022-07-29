package middleware

import (
	"encoding/json"
	"fmt"
	

	"CoinPrice_KryptoBackendTask/models" // models package where User schema is defined

	"github.com/gorilla/websocket"
)

//wss://stream.binance.com:9443/ws/BNBBTC@ticker

func AllAlerts(){
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var alerts []models.Alert

	// create the select sql query
	sqlStatement := `SELECT * FROM alerts`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)
	// unmarshal the row object to user

	// close the statement
	//defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var alert models.Alert

		// unmarshal the row object to user
		err = rows.Scan(&alert.ID, &alert.UserID, &alert.Symbol, &alert.Price)

		if err != nil {
			fmt.Print("Unable to scan the row. %v", err)
			return
		}

		alerts = append(alerts, alert)
	}

	for i:=0; i<len(alerts); i++{
		fmt.Println("Checking: "+alerts[i].Symbol+ " at "+alerts[i].Price)
		go getData(alerts[i].Symbol, alerts[i].Price, alerts[i].ID)
	}
}


func getData(symbol, price string, id int64) {
	
	// web socket is
	// websocket source
	c, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/!ticker@arr", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("Handshake Successful")

	}
	defer c.Close()
	// create the input channel
	inputStocks := make(chan []models.Trade)
	dogecoin := make(chan models.Trade)
	// producer: read from websocket and send to channel
	go func() {
		// read from the websocket
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				break
			}
			// unmarshal the message
			var trade []models.Trade
			json.Unmarshal(message, &trade)
			// send the trade to the channel
			inputStocks <- trade
		}
		close(inputStocks)
	}()
	// filter one kind of coin
	go func() {
		
			for trade := range inputStocks {
				for i := 0; i < len(trade); i++ {
					if trade[i].Symbol == symbol {
						dogecoin <- trade[i]
					}
				}

			}
		
		close(dogecoin)
	}()
	// print the trades
	for trade := range dogecoin {
		fmt.Println(trade.LastPrice)
		if trade.LastPrice <= price {
			fmt.Println("********* Price Found for alert  **********")
		}
		// json, _ := json.Marshal(trade)
		// fmt.Println(string(json))
	}

}
