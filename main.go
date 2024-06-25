package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a GET route with a parameter :name
	r.GET("/hello", func(c *gin.Context) {
		// Retrieve the value of the "name" parameter
		// name := c.Param("name")
		hostname, _ := os.Hostname()
		c.String(200, "Hello, %s!", hostname)
	})

	// Run the server on port 3000
	r.Run(":3000")
}
