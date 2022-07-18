package global

import (
	"github.com/i0Ek3/blogie/pkg/logger"
	"github.com/i0Ek3/blogie/pkg/setting"
)

var (
	// set global settings
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS

	// defined a logger object
	Logger *logger.Logger
)
