package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"

	"gortfoRio/bitflyer"

	// "GortfoRio/yahoo"
)
type JsonTestData struct {
	Name string
	Age int
}

func getXemHandler(w http.ResponseWriter, r *http.Request) {
	apiClient := bitflyer.New(os.Getenv("BITFLYER_API_KEY"), os.Getenv("BITFLYER_API_SECRET"))
	fmt.Println(apiClient.GetBalance())
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Headers","")
	testPerson1 := JsonTestData{"amano", 3}
	res, _ := json.Marshal(testPerson1)
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {

	http.HandleFunc("/xem", getXemHandler)
	http.ListenAndServe(":8080", nil)

	// fmt.Println(yahoo.GetYahooFinance)
}