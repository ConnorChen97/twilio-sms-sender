package main

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"

	"fmt"
)

func smsWebhook(c *gin.Context) {
	requestDump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func main() {
	router := gin.Default()
	router.POST("/twilio/sms/webhook", smsWebhook)

	router.Run("localhost:8080")
}
