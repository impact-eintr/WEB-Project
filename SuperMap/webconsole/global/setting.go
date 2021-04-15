package global

import (
	"database/sql"
	"webconsole/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	DB              *sql.DB
)
