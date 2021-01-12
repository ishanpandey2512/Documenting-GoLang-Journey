package main

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

type customer struct {
	Name    string
	Age     int
	Address string
}

func TestCustom(t *testing.T) {
	testcases := []struct {
		input  []byte
		output []byte
	}{
		{[]byte(`{"Name":"Ishan","Age":14,"Address":"India"}`), []byte(`not eligible`)},
		{[]byte(`{"Name":"Ishan","Address":"India"}`), []byte(`not eligible`)},
		{[]byte(`{"Name":"XYZ","Age":21,"Address":"DNFKLEF"}`), []byte(`{"Name":"XYZ","Age":21,"Address":"DNFKLEF"}`)},
	}

	t.Logf("Testing Customer Func Handler")
	for idx := range testcases {
		w := httptest.NewRecorder()
		// Method, Target, Convert to buffer before sending into body.
		req := httptest.NewRequest("POST", "http://localhost:8080/", bytes.NewBuffer(testcases[idx].input))
		handler(w, req)

		// w is the receiver argument, Result returns *http.Response
		resp := w.Result()
		// reads bytes from input output.
		body, _ := ioutil.ReadAll(resp.Body)
		// read the bytes, after converting them into type interface/ struct here, using unmarshall.
		// Create an output var customer
		//fmt.Println(string(body))
		if string(body) != string(testcases[idx].output) {
			t.Error("Failed")
			t.Logf("Expected: %s, \nGot %s", string(testcases[idx].output), string(body))
		}
	}
}
