package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new gin Router
	r := gin.Default()

	// Set the `requestHandler` function as a handler for every route
	r.NoRoute(requestHandler)

	// Start on port 3000
	r.Run(":3000")
}

// The Handler for every route
func requestHandler(c *gin.Context) {
	// Set the CORS headers
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")

	var url string
	uri := c.Request.RequestURI


	if strings.HasPrefix(uri, "/light") {
		url = "https://light.dsbcontrol.de"
		uri = strings.Replace(uri, "/light", "", 1)
	} else {
		url = "https://mobileapi.dsbcontrol.de"
			
	}

	// Send a request to the api with the URI
	res, err := http.Get(url+uri)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}
	defer res.Body.Close()

	// Read the body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		c.Status(500)
		return
	}

	// Send the status code of the request
	c.Status(res.StatusCode)
	// Send the body of the request
	fmt.Fprint(c.Writer, string(body))
}