package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/model"
	"github.com/i0Ek3/blogie/internal/routers"
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
	"github.com/i0Ek3/blogie/pkg/tracer"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port         string
	runMode      string
	config       string
	isVersion    bool
	buildTime    string
	buildVersion string
	gitCommitID  string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}

	err = setupSetting()
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

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
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

	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
		return
	}

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

func setupFlag() error {
	flag.StringVar(&port, "port", "", "run in which port")
	flag.StringVar(&runMode, "mode", "", "run in which mode")
	flag.StringVar(&config, "config", "configs/", "specify config path")
	flag.BoolVar(&isVersion, "version", false, "compile information")

	flag.Parse()
	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.AppSetting.ContextTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

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

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"blogie",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
