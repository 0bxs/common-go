package pgsql

import (
	"fmt"

	"github.com/0bxs/common-go/src/catch"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

type Pgsql struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Database     string `json:"database"`
	Schema       string `json:"schema"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	MaxIdleConns int    // 连接池的空闲数大小
	MaxOpenConns int    // 最大打开连接数
	ShowSql      bool
}

func Init(pgsql *Pgsql) *Engine {
	sourceName := fmt.Sprintf("user=%v password=%v dbname=%v search_path=%v host=%v port=%v sslmode=disable",
		pgsql.Username,
		pgsql.Password,
		pgsql.Database,
		pgsql.Schema,
		pgsql.Host,
		pgsql.Port,
	)
	engine := catch.Try1(xorm.NewEngine("pgx", sourceName))
	engine.ShowSQL(pgsql.ShowSql)
	engine.SetMaxIdleConns(pgsql.MaxIdleConns)
	engine.SetMaxOpenConns(pgsql.MaxOpenConns)
	engine.SetMapper(names.GonicMapper{})
	catch.Try(engine.Ping())
	return new(engine)
}
