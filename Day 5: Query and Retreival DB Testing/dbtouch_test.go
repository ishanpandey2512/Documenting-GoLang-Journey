package main

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateCustomer(t *testing.T) {

	db, err := sql.Open("mysql", "root:password@/customer_service?multiStatements=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	q1 := "DROP TABLE Address"
	q2 := "DROP TABLE  Customer"
	q3 := "CREATE TABLE Customer(ID INT NOT NULL auto_increment, Name varchar(255), DOB varchar(255), PRIMARY KEY(ID))"
	q4 := "CREATE TABLE Address(ID INT NOT NULL auto_increment, StreetName varchar(255), City varchar(255), State varchar(255),Cus_ID INT , PRIMARY KEY (ID), FOREIGN KEY (Cus_ID) REFERENCES Customer(ID))"

	_, err = db.Query(q1)
	_, err = db.Query(q2)
	_, err = db.Query(q3)
	_, err = db.Query(q4)

	if err != nil {
	}
	testCases := []struct {
		inp Customer
		out Customer
	}{
		{Customer{ID: 1, Name: "Ishan Pandey", DOB: "25-12-1999", Addr: Address{ID: 1, StreetName: "SarvaSiddhi Vinayak Colony", City: "Lucknow", State: "UP", Cus_ID: 1}},
			Customer{ID: 1, Name: "Ishan Pandey", DOB: "25-12-1999", Addr: Address{ID: 1, StreetName: "SarvaSiddhi Vinayak Colony", City: "Lucknow", State: "UP", Cus_ID: 1}}},

		{Customer{ID: 2, Name: "Varun Singh", DOB: "15-02-1996", Addr: Address{ID: 2, StreetName: "LDA Colony, Sector-I", City: "Lucknow", State: "UP", Cus_ID: 2}},
			Customer{ID: 2, Name: "Varun Singh", DOB: "15-02-1996", Addr: Address{ID: 2, StreetName: "LDA Colony, Sector-I", City: "Lucknow", State: "UP", Cus_ID: 2}}},
	}

	for i := range testCases {
		d := CreateCustomer(db, testCases[i].inp)
		if !reflect.DeepEqual(d, testCases[i].out) {
			t.Errorf("FAILED for %v expected %v got %v\n", testCases[i].inp, testCases[i].out, d)
		}
	}

}
