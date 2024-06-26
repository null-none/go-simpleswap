package simpleswap

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetExchange(apiKey string, limit int, offset int) string {
	requestURL := fmt.Sprintf("https://api.simpleswap.io/get_exchange?api_key=%s&%d&%d", apiKey, limit, offset)
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
