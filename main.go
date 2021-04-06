package main

import (
	"fmt"
	"os"

	"gortfoRio/bitflyer"
)

func main() {
	apiClient := bitflyer.New(os.Getenv("BITFLYER_API_KEY"), os.Getenv("BITFLYER_API_SECRET"))
	fmt.Println(apiClient.GetBalance())
}