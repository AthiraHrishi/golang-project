/*Using GIN Framework demonstrate custom binding request(At least give 5 field example with
including boolean, pointers, int, string, runes.*/

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CustomRequest struct {
	Name    string  `json:"name" binding:"required"`             // String
	Age     int     `json:"age" binding:"gte=0"`                 // Integer, should be greater than or equal to 0
	Active  bool    `json:"active"`                              // Boolean
	Comment *string `json:"comment"`                             // Pointer to a string
	Letter  rune    `json:"letter" binding:"len=1"`              // Rune, should be exactly one character
}

func CustomBindingHandler(c *gin.Context) {
	var req CustomRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name":    req.Name,
		"Age":     req.Age,
		"Active":  req.Active,
		"Comment": req.Comment,
		"Letter":  string(req.Letter),
	})
}

func main() {
	r := gin.Default()

	r.POST("/custom-bind", CustomBindingHandler)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
