package main

import (
	"database/sql"
	"fmt"
	"log"
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

func CreateCustomer(db *sql.DB, c Customer) Customer {
	query := `insert into Customer values (?,?,?)`

	var custData []interface{}

	custData = append(custData, c.ID)
	custData = append(custData, c.Name)
	custData = append(custData, c.DOB)

	_, err := db.Query(query, custData...)

	query = `insert into Address values (?,?,?,?,?)`

	var addData []interface{}

	addData = append(addData, c.Addr.ID)
	addData = append(addData, c.Addr.StreetName)
	addData = append(addData, c.Addr.City)
	addData = append(addData, c.Addr.State)
	addData = append(addData, c.Addr.Cus_ID)

	_, err = db.Query(query, addData...)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	var rows *sql.Rows
	rows, err = db.Query("SELECT * FROM Customer inner join Address on Address.Cus_id=Customer.id and Customer.id= ? order by Customer.id asc, Address.id asc", c.ID)

	if err != nil {
		fmt.Println("Error is ", err)
	}

	var customer Customer
	for rows.Next() {
		if err = rows.Scan(&customer.ID, &customer.Name, &customer.DOB, &customer.Addr.ID, &customer.Addr.StreetName, &customer.Addr.City, &customer.Addr.State, &customer.Addr.Cus_ID); err != nil {
			log.Fatal(err)
		}
	}
	return customer
}
