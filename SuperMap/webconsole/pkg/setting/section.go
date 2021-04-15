package setting

import (
	"log"
)

type ServerSettingS struct {
	Port int
}

type CacheSettingS struct {
	CacheType string
	Port      string
	TTL       int
	CacheDir  string
}

type DatabaseSettingS struct {
	Host     string
	Port     int
	User     string
	Password string
	DBname   string
}

func (s *Setting) ReadSection(key string, v interface{}) error {
	err := s.vp.UnmarshalKey(key, v)
	if err != nil {
		log.Println(key)
		return err
	}
	return nil
}
