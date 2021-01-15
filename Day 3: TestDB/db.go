package main

import (
	"database/sql"
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

// func getCustomer returns a slice of type Customer as result, for a SQL query into database for id.
func getCustomer(db *sql.DB, id int) []Customer {
	var ids []interface{}
	// When ID == 0, does not exist.
	query := `SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID ORDER BY Customer.ID, Address.Cus_ID;`
	// If id is given.
	if id != 0 {
		query = `SELECT * FROM Customer INNER JOIN Address ON Customer.ID = Address.Cus_ID WHERE Customer.ID = ? ORDER BY Customer.ID, Address.Cus_ID; `
		ids = append(ids, id)
	}

	rows, err := db.Query(query, ids...)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var res []Customer
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.ID, &c.Name, &c.DOB, &c.Addr.ID, &c.Addr.StreetName, &c.Addr.City, &c.Addr.State, &c.Addr.Cus_ID); err != nil {
			log.Fatal(err)
		}
		res = append(res, c)
	}
	return res
}
