package configs

type Configs struct {
	Log     LogConfig
	DB      DBConfig
	Project ProjectConfig
}

// LogConfig 日志配置
type LogConfig struct {
	DevMode bool   // 模式，如果是dev则日志打印到控制台，否则输出到文件
	Level   string `yaml:"level"`
	Path    string `yaml:"path"`
	Save    uint   `yaml:"save"` // 最大文件保存份数
}

// DBConfig 数据库配置
type DBConfig struct {
	DBType          string // 数据库类型
	Username        string // 用户名
	Password        string // 密码
	Host            string // host
	Port            int    // 端口
	DB              string // 数据库
	Param           string // 额外参数
	ShowSql         bool   // 是否显示sql
	MaxIdleConns    int    // 最大空闲连接数
	MaxOpenConns    int    // 最大打开连接数
	ConnMaxLifetime int    // 连接最大生存周期,单位：s
}

// ProjectConfig 项目配置
type ProjectConfig struct {
	ServerPort      string
	ExcludeAuthUrls []string
	AllowMaxDelay   int // 最大允许延迟，单位：s
}


