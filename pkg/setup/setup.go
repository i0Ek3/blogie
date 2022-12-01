package setup

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/i0Ek3/blogie/global"
	"github.com/i0Ek3/blogie/internal/model"
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
	"github.com/i0Ek3/blogie/pkg/tracer"
	"github.com/i0Ek3/blogie/pkg/validator"
	"github.com/i0Ek3/blogie/pkg/version"
	"github.com/natefinch/lumberjack"
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
	global.GDB, err = model.NewDBEngine(global.DatabaseSetting)
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
	global.Trans = uni

	return nil
}
