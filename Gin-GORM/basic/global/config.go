package global

type Config struct {
	Port        int         `mapstructrue:"port"`
	CacheConfig CacheConfig `mapstructrue:"cache"`
	MysqlConfig MysqlConfig `mapstructrue:"mysql"`
}

type MysqlConfig struct {
	Host     string `mapstructrue:"host"`
	Port     int    `mapstructrue:"port"`
	User     string `mapstructrue:"user"`
	Password string `mapstructrue:"password"`
	DBname   string `mapstructrue:"dbname"`
}

type CacheConfig struct {
	CacheType string `mapstructrue:"cachetype"`
	Port      int    `mapstructrue:"port"`
	TTL       int    `mapstructrue:"tll"`
	CacheDir  string `mapstructrue:"cacheDir"`
}

// 配置修改后如何不关停？
var G = new(Config)
