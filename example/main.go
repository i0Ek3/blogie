package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// gin.Default() -> gin.New() -> Logger/Recovery -> r.GET() -> r.Run()
	// gin.Default() used to create an Engine instance which import Logger and Recovery middleware.
	// gin.New() initializes Engine instance and return.
	r := gin.Default()

	// r.GET() registers /ping router into handler.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// uncomment it to run
	// global.Logger.Infof(context.Background(), "This is a test message to test Info level.")

	// r.Run() parses the given address and then invoke http.ListenAndServe() register
	// an Engine instance into handler, also Engine type implements ServeHTTP(), so Engine
	// can be passed by a parameter.
	_ = r.Run()
}
