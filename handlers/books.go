package handlers

import (
	"github.com/andey-robins/bookshop-go/db"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

type Book struct {
	Id     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

func CreateBook(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 3); err != nil {
		log.Println("In POST books/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("Title", jsonRaw["Title"]); err != nil {
		log.Println("In POST books/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("Author", jsonRaw["Author"]); err != nil {
		log.Println("In POST books/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidatePositiveNumber("Price", jsonRaw["Price"]); err != nil {
		log.Println("In POST books/new, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Book

	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}



	_, err := db.CreateBook(json.Title, json.Author, json.Price)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"status": "success"})
}

func GetPrice(c *gin.Context) {
	jsonRaw := make(map[string]interface{})

	if err := c.ShouldBindBodyWith(&jsonRaw, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateJsonLength(jsonRaw, 2); err != nil {
		log.Println("In GET books/price, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("Title", jsonRaw["Title"]); err != nil {
		log.Println("In GET books/price, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := ValidateNonEmptyString("Author", jsonRaw["Author"]); err != nil {
		log.Println("In GET books/price, " + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var json Book
	if err := c.ShouldBindBodyWith(&json, binding.JSON); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	bid, err := db.GetBookId(json.Title, json.Author)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	price, err := db.GetBookPrice(bid)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"price": price})
}
