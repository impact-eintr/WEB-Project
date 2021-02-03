package setting

//配置文件的对应结构体
import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	LogFileName           string
	//UploadServerUrl string
	//UploadImageMaxSize int
	//UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (this *Setting) ReadSection(k string, v interface{}) error {
	err := this.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}
