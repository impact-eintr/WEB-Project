package global

import (
	"database/sql"
	"webconsole/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	LoggerSetting   *setting.LoggerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	DB              *sql.DB
)
