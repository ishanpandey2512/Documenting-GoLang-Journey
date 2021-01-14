package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	db *sql.DB
}

var (
	ctx context.Context
)

type Customer struct {
	ID   int
	Name string
	DOB  string
	Addr Address
}

type Address struct {
	ID         int
	StreetName string
	City       string
	State      string
	Cus_ID     int
}

/*
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "ishan"
	dbPass := "Tocmsoon25"
	dbName := "customer_service"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
*/

func getCustomerAll(db *sql.DB, id int) []Customer {

	// Establishing connection with the database.

	/*
		db, err := sql.Open("mysql", "root:password@/customer_service")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	*/

	// When ID == 0, does not exist.
	if id == 0 {
		stmtOut, err := db.Query(`SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID ORDER BY Customer.ID, Address.Cus_ID;`)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtOut.Close()
		var res []Customer
		for stmtOut.Next() {

			var c Customer
			if err := stmtOut.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.Cus_ID); err != nil {
				log.Fatal(err)
			}
			//fmt.Println(c.ID, c.Name, c.DOB, c.Addr.ID, c.Addr.StreetName, c.Addr.City, c.Addr.State, c.Addr.Cus_ID)
			res = append(res, c)
		}

		return res
	} else {
		// Else when ID is present/ not found in DB.
		stmtOut, err := db.Query(fmt.Sprintf(`SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID WHERE Customer.ID = %v ORDER BY Customer.ID, Address.Cus_ID; `, id))
		//fmt.Println(err)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtOut.Close()

		var res []Customer
		for stmtOut.Next() {
			var c Customer
			if err := stmtOut.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.Cus_ID); err != nil {
				log.Fatal(err)
			}
			//fmt.Println(c.ID, c.Name, c.DOB, c.Addr.ID, c.Addr.StreetName, c.Addr.City, c.Addr.State, c.Addr.Cus_ID)
			res = append(res, c)
		}
		return res
	}
}

func main() {

	/*
		// Establishing connection with the database.
		db, err := sql.Open("mysql", "root:password@/customer_service")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	*/

	/*
		    For all the results having a particular ID.
			res := getCustomerAll(db, 1)
			for i := range res {
				fmt.Println(res[i])
			}

			return res
	*/

	/*
		Sending data to the server.
		insert, err2 := db.Prepare("INSERT INTO Customer VALUES( ?, ?, ?)")
		_, err2 = insert.Exec(4, "Paresh Rawal", "15-02-1972")
		if err2 != nil {
			panic(err2.Error())
		}
	*/

}
