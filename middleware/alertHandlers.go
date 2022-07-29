package middleware

import (
	
	"database/sql"
	"encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"
	"log"
	"net/http" // used to access the request and response object of the api

	"strconv"  // package used to covert string into int type
	

	"CoinPrice_KryptoBackendTask/models" // models package where User schema is defined



	"github.com/gorilla/mux" // used to get the params from the route

	_ "github.com/lib/pq"         // postgres golang driver
)

// CreateAlert create a alert in the postgres db
func CreateAlert(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty user of type models.User
	var alert models.Alert

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&alert)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	

	// call insert user function and pass the user
	insertID := insertAlert(alert)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Alert created successfully",
	}

	// send the response
	go AllAlerts()
	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its id
func GetAlert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	alert, err := getAlert(int64(id))

	if err != nil {
		log.Fatalf("Unable to get alert. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(alert)
}

// GetAllUser will return all the users
func GetAllUserAlerts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the users in the db
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	alerts, err := getAllAlerts(int64(id))

	if err != nil {
		log.Fatalf("Unable to get all alerts. %v", err)
	}

	// send all the users as response
	json.NewEncoder(w).Encode(alerts)
}

// UpdateUser update user's detail in the postgres db
func UpdateAlert(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty user of type models.User
	var alert models.Alert

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&alert)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update user to update the user
	updatedRows := updateAlert(int64(id), alert)

	// format the message string
	msg := fmt.Sprintf("Alert updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteUser delete user's detail in the postgres db
func DeleteAlert(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteUser, convert the int to int64
	deletedRows := deleteAlert(int64(id))

	// format the message string
	msg := fmt.Sprintf("Alert updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one user in the DB
func insertAlert(alert models.Alert) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning userid will return the id of the inserted user
	sqlStatement := `INSERT INTO alerts (userid, symbol, price) VALUES ($1, $2, $3) RETURNING alertid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, alert.UserID, alert.Symbol, alert.Price).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one user from the DB by its userid
func getAlert(id int64) (models.Alert, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a user of models.User type
	var alert models.Alert

	// create the select sql query
	sqlStatement := `SELECT * FROM alerts WHERE alertid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to user
	err := row.Scan(&alert.ID, &alert.UserID, &alert.Symbol, &alert.Price)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return alert, nil
	case nil:
		return alert, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return alert, err
}

// get one user from the DB by its userid
func getAllAlerts(id int64) ([]models.Alert, error) {
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
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		if(alert.UserID == id){
			alerts = append(alerts, alert)
		}
			
	}

	// return empty user on error
	return alerts, err
}

// update user in the DB
func updateAlert(id int64, alert models.Alert) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE alerts SET symbol=$2, price=$3 WHERE alertid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, alert.Symbol, alert.Price)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete user in the DB
func deleteAlert(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM alerts WHERE alertid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}