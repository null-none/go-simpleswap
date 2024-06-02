package main

import (
	"net/http"

	Simpleswap "github.com/null-none/go-simpleswap/simpleswap"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	apiKey := ""

	r.GET("/get_all_currencies", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"get_all_currencies": Simpleswap.GetAllCurrencies(apiKey),
		})
	})

	r.GET("/get_currency", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"get_currency": Simpleswap.GetCurrency(apiKey),
		})
	})

	r.GET("/get_market_info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"get_currency": Simpleswap.GetMarketInfo(),
		})
	})

	r.GET("/get_exchange", func(c *gin.Context) {
		id := "btc"
		c.JSON(http.StatusOK, gin.H{
			"get_currency": Simpleswap.GetExchange(apiKey, id),
		})
	})

	r.Run()
}
