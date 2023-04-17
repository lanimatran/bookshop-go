package handlers

import (
	"fmt"
	"net/http"

	"github.com/andey-robins/bookshop-go/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

type Order struct {
	CustomerId int  `json:"customerId"`
	BookId     int  `json:"bookId"`
	Shipped    bool `json:"shipped"`
}

func CreateOrder(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In POST orders/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("BookId", jsonRaw["BookId"]); err != nil {
		log.Println("In POST orders/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("CustomerId", jsonRaw["CustomerId"]); err != nil {
		log.Println("In POST orders/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Order
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.CreatePO(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func GetShipmentStatus(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In GET orders/shipped, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("BookId", jsonRaw["BookId"]); err != nil {
		log.Println("In GET orders/shipped, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("CustomerId", jsonRaw["CustomerId"]); err != nil {
		log.Println("In GET orders/shipped, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Order
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	status, err := db.IsPOShipped(pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": status})
}

func ShipOrder(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In PUT orders/ship, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("BookId", jsonRaw["BookId"]); err != nil {
		log.Println("In PUT orders/ship, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("CustomerId", jsonRaw["CustomerId"]); err != nil {
		log.Println("In PUT orders/ship, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Order
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = db.ShipPO(pid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func GetOrderStatus(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In GET orders/status, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("BookId", jsonRaw["BookId"]); err != nil {
		log.Println("In GET orders/status, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("CustomerId", jsonRaw["CustomerId"]); err != nil {
		log.Println("In GET orders/status, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Order
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	pid, err := db.GetPOByContents(json.BookId, json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	addr, err := db.GetCustomerAddress(json.CustomerId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<html>
		<head>
		<title>Order Status</title>
		</head>
		<body>
		<h1>Order Status</h1>
		<p>Order ID: `+fmt.Sprint(pid)+`</p>
		<p>Book ID: `+fmt.Sprint(json.BookId)+`</p>
		<p>Customer ID: `+fmt.Sprint(json.CustomerId)+`</p>
		<p>Shipping Address: `+fmt.Sprint(addr)+`</p>
		</body>
	</html>`))
}
