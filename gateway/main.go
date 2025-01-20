package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Capture all methods/paths
	r.Any("/*proxyPath", func(c *gin.Context) {
		path := c.Param("proxyPath")
		if path == "" || path == "/" {
			c.String(http.StatusNotFound, "Endpoint not found in API Gateway")
			return
		}

		// Determine the target service based on path prefix
		var targetURL string
		switch {
		case strings.HasPrefix(path, "/service1"):
			targetURL = "http://localhost:9001" + strings.TrimPrefix(path, "/service1")
		case strings.HasPrefix(path, "/service2"):
			targetURL = "http://localhost:9002" + strings.TrimPrefix(path, "/service2")
		case strings.HasPrefix(path, "/service3"):
			targetURL = "http://localhost:9003" + strings.TrimPrefix(path, "/service3")
		default:
			c.String(http.StatusNotFound, "Endpoint not found in API Gateway")
			return
		}

		// Create a new request with the same method and body as the incoming request
		req, err := http.NewRequest(c.Request.Method, targetURL, c.Request.Body)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to create request to service")
			return
		}

		// Copy all headers
		for k, v := range c.Request.Header {
			req.Header[k] = v
		}

		// Send the request to the target service
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.String(http.StatusBadGateway, "Failed to call service")
			return
		}
		defer resp.Body.Close()

		// Copy response headers
		for k, v := range resp.Header {
			c.Writer.Header()[k] = v
		}
		// Set the status code to the one from the service
		c.Status(resp.StatusCode)

		// Copy the response body to the gateway response
		io.Copy(c.Writer, resp.Body)
	})

	log.Println("API Gateway is running on port 8080")
	r.Run(":8080")
}
