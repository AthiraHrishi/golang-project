//Create a post api which will have rate limit on basis of each and unique ip.

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	middleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memory "github.com/ulule/limiter/v3/drivers/store/memory"
	"net/http"
	"time"
)

func main() {
	// Create a new rate limiter with a limit of 5 requests per minute per IP
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  5,
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)

	// Initialize Gin router
	router := gin.Default()

	// Apply the rate limiter middleware to all routes
	router.Use(middleware.NewMiddleware(instance))

	// Define a POST endpoint
	router.POST("/submit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request successful",
		})
	})

	// Start the server
	router.Run(":8080")
}
