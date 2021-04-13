package setting

import "log"

type ServerSettingS struct {
	Port int
}

type CacheSettingS struct {
	CacheType string
	Port      int
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
	log.Println(key, v)
	if err != nil {
		return err
	}
	return nil
}
