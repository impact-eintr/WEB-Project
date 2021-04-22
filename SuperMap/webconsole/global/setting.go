package global

import (
	"database/sql"
	"webconsole/pkg/setting"
)

var (
	Conf            *setting.Setting
	ServerSetting   *setting.ServerSettingS
	LoggerSetting   *setting.LoggerSettingS
	DatabaseSetting *setting.DatabaseSettingS
	CacheSetting    *setting.CacheSettingS
	DB              *sql.DB
)

func Init() (err error) {
	Conf, err = setting.NewSetting()
	if err != nil {
		return
	}
	_ = Conf
	return

}
