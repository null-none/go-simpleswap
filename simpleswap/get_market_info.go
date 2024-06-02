package simpleswap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetMarketInfo(apiKey string) string {
	requestURL := fmt.Sprintf("https://api.simpleswap.io/get_market_info?api_key=%s&%s", apiKey)
	response, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body: %s\n", err)
		os.Exit(1)
	}
	return string(body)
}
