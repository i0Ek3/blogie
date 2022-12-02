package global

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
	"github.com/i0Ek3/blogie/pkg/validator"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS
	EnableSetting   *setting.EnableSettingS
	RedisSetting    *setting.RedisSettingS
	Logger          *logger.Logger
)

var (
	// Global DB
	GDB *gorm.DB

	// Global Validator
	Validator *validator.CustomValidator
	Trans     *ut.UniversalTranslator

	// Global Tracer
	Tracer opentracing.Tracer
)
