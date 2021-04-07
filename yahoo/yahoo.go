package yahoo

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetYahooFinance() {

	url := "https://yahoo-finance15.p.rapidapi.com/api/yahoo/ne/news"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "...")
	req.Header.Add("x-rapidapi-host", "yahoo-finance15.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}