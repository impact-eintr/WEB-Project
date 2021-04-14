package global

import (
	"basic/pkg/setting"
	"database/sql"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	DB              *sql.DB
)
