package main

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	_a, _b, _c := false, false, true
	if _a {
		runme()
	}
	if _b {
		watchdog()
	}
	if _c {
		decode()
	}
}

func runme() {
	// gin.Default() -> gin.New() -> Logger/Recovery -> r.GET() -> r.Run()
	// gin.Default() used to create an Engine instance which import Logger and Recovery middleware.
	// gin.New() initializes Engine instance and return.
	r := gin.Default()

	// r.GET() registers /ping router into handler.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Test log output
	global.Logger.Infof(context.Background(), "This is a test message to test Info level.")

	// r.Run() parses the given address and then invoke http.ListenAndServe() register
	// an Engine instance into handler, also Engine type implements ServeHTTP(), so Engine
	// can be passed by a parameter.
	_ = r.Run()
}

func watchdog() {
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	path := "/path/to/config.yaml"
	_ = watcher.Add(path)
	<-done
}

func decode() {
	// Payload fetched from GenerateToken(), which contains appKey and appSecret
	// after somebody gets your payload, your appKey and appSecret will be cracked
	// by following method, so please do not store plaintext information in the payload
	payload := "eyJhcHBfa2V5IjoiY2UwMTM2ZWJiZmU5MzgzZWM4ZjM1YTRlNjFiNmM2NjciLCJhcHBfc2VjcmV0IjoiYjVkZGU2M2U3OWQ5MmRhMjUwMmM5YTMxNjBhNWY2NTUiLCJpc3MiOiJibG9naWUifQ"
	msg, _ := base64.StdEncoding.DecodeString(payload)
	log.Println(string(msg))
}
