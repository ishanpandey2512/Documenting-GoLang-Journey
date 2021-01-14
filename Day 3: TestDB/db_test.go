package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCustom(t *testing.T) {
	// Input: ID, Output: expected complete result from MYSQL server.
	testcases := []struct {
		input  int
		output []Customer
	}{
		{1, []Customer{
			{1, "Ishan Pandey", "25-12-1999", Address{1, "SarvaSiddhi Vinayak Colony, Vrindavan", "Lucknow", "UP", 1}},
			{1, "Ishan Pandey", "25-12-1999", Address{3, "St. Peters Church", "Moscow", "24th State, Russia", 1}},
		}},

		{2, []Customer{{2, "Errichto", "7-11-1993", Address{2, "151/1, 5th Cross, Koramangla", "Bangalore", "Karnataka", 2}}}},

		// when no data is given.
		{0, []Customer{
			{1, "Ishan Pandey", "25-12-1999", Address{1, "SarvaSiddhi Vinayak Colony, Vrindavan", "Lucknow", "UP", 1}},
			{1, "Ishan Pandey", "25-12-1999", Address{3, "St. Peters Church", "Moscow", "24th State, Russia", 1}},
			{2, "Errichto", "7-11-1993", Address{2, "151/1, 5th Cross, Koramangla", "Bangalore", "Karnataka", 2}},
			{12, "Darisa Manish Kumar", "12-10-1999", Address{4, "Some Address", "Tiruvalur", "Karnataka", 12}},
		}},

		{12, []Customer{{12, "Darisa Manish Kumar", "12-10-1999", Address{4, "Some Address", "Tiruvalur", "Karnataka", 12}}}},
	}

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

	for idx := range testcases {
		actual := getCustomerAll(db, testcases[idx].input)
		if cmp.Equal(actual, testcases[idx].output) == false {
			t.Error("Failed")
			t.Logf("Expected: %v, \n Got: %v", testcases[idx].output, actual)
			fmt.Println(actual)
		}
	}

}
