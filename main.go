package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/model"
	"github.com/i0Ek3/blogie/internal/routers"
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, // maximum size, megabytes
		MaxAge:    10,  // retain days
		LocalTime: true,
	}, "", log.LstdFlags)

	return nil
}

// @title blogie
// @version 1.0
// @description A blog backend program developed with Gin.
//  @termOfService https://github.com/i0Ek3/blogie
func main() {

	/* Gin demo
	   // gin.Default() -> gin.New() -> Logger/Recovery -> r.GET() -> r.Run()
	   // gin.Default() use to create an Engine instance which import Logger and Recovery middleware.
	   // gin.New() initializes Engine instance and return.
	   r := gin.Default()

	   // r.GET() registers /ping router into handler.
	   r.GET("/ping", func(c *gin.Context) {
	       c.JSON(200, gin.H{"message": "pong"})
	   })

	   // r.Run() parses the given address and then invoke http.ListenAndServe() register
	   // an Engine instance into handler, also Engine type implements ServeHTTP(), so Engine
	   // can be passed by a parameter.
	   r.Run()
	*/

	// set run mode for Gin
	gin.SetMode(global.ServerSetting.RunMode)

	// create an Engine instance which is a handler
	router := routers.NewRouter()

	// create a http server by our own rules
	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// uncomment it to run
	// global.Logger.Infof("This is a test message to test Info level.")

	ser.ListenAndServe()
}