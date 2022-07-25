package global

import (
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
)

var (
	// ServerSetting set global settings
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	JWTSetting      *setting.JWTSettingS
	EmailSetting    *setting.EmailSettingS

	// Logger defines a logger object
	Logger *logger.Logger
)
