package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/model"
	"github.com/i0Ek3/blogie/internal/routers"
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
	"github.com/i0Ek3/blogie/pkg/tracer"
	"github.com/i0Ek3/blogie/pkg/validator"
	"github.com/i0Ek3/blogie/pkg/version"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	port    string
	runMode string
	config  string
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

	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}
}

// @title blogie
// @version 1.0
// @description A blog backend program developed with Gin.
// @termOfService https://github.com/i0Ek3/blogie
func main() {
	// Set run mode for Gin
	gin.SetMode(global.ServerSetting.RunMode)

	// Create an Engine instance which is a handler
	router := routers.NewRouter()

	// Create a http server by our own rules
	ser := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: global.ServerSetting.HeaderBytes,
	}

	go func() {
		err := ser.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("main::ser.ListenAndServe err: %v", err)
		}
	}()

	quit := make(chan os.Signal, 2)
	// NOTES: 2 SIGINT, 15 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := ser.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting...")
}

func setupFlag() error {
	var (
		isVersion    bool
		buildTime    string
		buildVersion string
		gitCommitID  string
	)

	if isVersion {
		fmt.Printf("build_time: %s\n", buildTime)
		fmt.Printf("build_version: %s\n", buildVersion)
		fmt.Printf("git_commit_id: %s\n", gitCommitID)
	}

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
	//fmt.Println("DEBUG------->global.ServerSetting", global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	//fmt.Println("DEBUG------->global.AppSetting", global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	//fmt.Println("DEBUG------->global.DatabaseSetting", global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("JWT", &global.JWTSetting)
	//fmt.Println("DEBUG------->global.JWTSetting", global.JWTSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Email", &global.EmailSetting)
	//fmt.Println("DEBUG------->global.EmailSetting", global.EmailSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Enable", &global.EnableSetting)
	//fmt.Println("DEBUG------->global.EnableSetting", global.EnableSetting)
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
		MaxSize:   100, // maximum size, megabytes
		MaxAge:    7,   // retain days
		LocalTime: true,
	}, "", log.LstdFlags)

	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		version.AppName,
		version.Address,
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator
	uni := ut.New(en.New(), zh.New())
	v, ok := binding.Validator.Engine().(*v10.Validate)
	if ok {
		zhTran, _ := uni.GetTranslator("zh")
		enTran, _ := uni.GetTranslator("en")
		err := zh_translations.RegisterDefaultTranslations(v, zhTran)
		if err != nil {
			return err
		}
		err = en_translations.RegisterDefaultTranslations(v, enTran)
		if err != nil {
			return err
		}
	}
	global.Ut = uni
	return nil
}
