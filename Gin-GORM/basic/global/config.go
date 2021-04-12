package global

type Config struct {
	Port            int             `mapstructrue:"port"`
	MysqlConfig     MysqlConfig     `mapstructrue:"mysql"`
	MemCacheConfig  MemCacheConfig  `mapstructrue:"memcache"`
	DiskCacheConfig DiskCacheConfig `mapstructrue:"diskcache"`
}

type MysqlConfig struct {
	Host     string `mapstructrue:"host"`
	Port     int    `mapstructrue:"port"`
	User     string `mapstructrue:"user"`
	Password string `mapstructrue:"password"`
	DBname   string `mapstructrue:"dbname"`
}

type MemCacheConfig struct {
	Port int `mapstructrue:"port"`
	TTL  int `mapstructrue:"tll"`
}

type DiskCacheConfig struct {
	Port     int    `mapstructrue:"port"`
	TTL      int    `mapstructrue:"tll"`
	CacheDir string `mapstructrue:"cacheDir"`
}

// 配置修改后如何不关停？
