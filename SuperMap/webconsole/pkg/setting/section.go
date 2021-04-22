package setting

import (
	"log"
)

type ServerSettingS struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Locale  string `mapstructure:"locale"`
	Port    int    `mapstructure:"port"`

	StartTime int   `mapstructure:"starttime"`
	MachineID int64 `mapstructure:"machine_id"`
}

type LoggerSettingS struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_bacups"`
}
type CacheSettingS struct {
	CacheType string `mapstructure:"cachetype"`
	Port      string `mapstructure:"port"`
	TTL       int    `mapstructure:"ttl"`
	CacheDir  string `mapstructure:"cachedir"`
}

type DatabaseSettingS struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DBname       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idel_conns"`
}

func (s *Setting) ReadSection(key string, v interface{}) error {
	err := s.vp.UnmarshalKey(key, v)
	if err != nil {
		log.Println(key)
		return err
	}
	return nil
}
