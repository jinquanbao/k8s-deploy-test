package inits

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"kgdt/configs"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

// InitDB 初始化数据库
func InitDB(config configs.DBConfig) *xorm.Engine{
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.Username, config.Password, config.Host, config.Port, config.DB, config.Param)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	engine, err := xorm.NewEngine(config.DBType, dataSourceName)
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败:%s", err))
	}

	engine.SetLogger(xlog.NewSimpleLogger(LogWriter))
	engine.ShowSQL(config.ShowSql)
	engine.SetMaxIdleConns(config.MaxIdleConns)
	engine.SetMaxOpenConns(config.MaxOpenConns)
	engine.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)
	err = engine.Ping()
	if err != nil {
		panic(fmt.Sprintf("数据库初始化失败:%s", err))
	}
	return engine
}
