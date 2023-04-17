package handlers

import (
	"github.com/andey-robins/bookshop-go/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

type Customer struct {
	Id             int     `json:"id"`
	Name           string  `json:"name"`
	ShippingAddr   string  `json:"shippingAddr"`
	AccountBalance float32 `json:"accountBalance"`
}

func CreateCustomer(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In POST customers/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("Name", jsonRaw["Name"]); err != nil {
		log.Println("In POST customers/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("ShippingAddr", jsonRaw["ShippingAddr"]); err != nil {
		log.Println("In POST customers/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Customer
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := db.CreateCustomer(json.Name, json.ShippingAddr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func UpdateCustomerAddress(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In PUT customers/updateAddress, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("ShippingAddr", jsonRaw["ShippingAddr"]); err != nil {
		log.Println("In PUT customers/updateAddress, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("Id", jsonRaw["Id"]); err != nil {
		log.Println("In PUT customers/updateAddress, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Customer
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := db.UpdateCustomerAddress(json.Id, json.ShippingAddr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"status": "success"})
}

func GetCustomerBalance(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 1); err != nil {
		log.Println("In GET customers/balance, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("Id", jsonRaw["Id"]); err != nil {
		log.Println("In GET customers/balance, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Customer
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	balance, err := db.CustomerBalance(json.Id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"balance": balance})
}
