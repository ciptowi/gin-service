package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type contact struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var contacts = []contact{
	{ID: "1", Name: "name 1", Phone: "1231212412322", Email: "email@email.com"},
	{ID: "2", Name: "name 2", Phone: "1231212412322", Email: "email@email.com"},
	{ID: "3", Name: "name 3", Phone: "1231212412322", Email: "email@email.com"},
}

func getContacts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, contacts)
}

func postContacts(c *gin.Context) {
	var newContact contact

	if err := c.BindJSON(&newContact); err != nil {
		return
	}

	contacts = append(contacts, newContact)
	c.IndentedJSON(http.StatusCreated, newContact)
}

func getContactByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range contacts {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "contact not found"})
}

func main() {
	router := gin.Default()

	router.POST("/contact", postContacts)
	router.GET("/contact", getContacts)
	router.GET("/contact/:id", getContactByID)

	router.Run("localhost:8080")
}
