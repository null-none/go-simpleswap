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

	r.GET("/get_exchanges", func(c *gin.Context) {
		type Validate struct {
			Limit  int `form:"limit" json:"limit" xml:"limit"  binding:"required"`
			Offset int `form:"offest" json:"offest" xml:"offest"  binding:"required"`
		}
		var json Validate
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"get_currency": Simpleswap.GetExchange(apiKey, json.Limit, json.Offset),
		})
	})

	r.GET("/get_exchange", func(c *gin.Context) {
		type Validate struct {
			Id string `form:"id" json:"id" xml:"id"  binding:"required"`
		}
		var json Validate
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"get_currency": Simpleswap.GetExchanges(apiKey, json.Id),
		})
	})

	r.Run()
}
