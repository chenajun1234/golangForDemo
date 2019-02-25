package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/lunny/log"
	"xormDemo/model"
)

//引擎
var XEngine *xorm.Engine

func init() {
	var err error
	XEngine, err = xorm.NewEngine("mysql", "root:root@/gotest?charset=utf8mb4")
	defer XEngine.Close()
	if err != nil {
		log.Printf("链接mysql异常 %v", err)
	}
	XEngine.ShowSQL(true)
	err = XEngine.Sync(new(model.User))
	if err != nil {
		log.Printf("同步表User异常 %v", err)
	}

}

func InsertUer(u model.User) {
	i, e := XEngine.Insert(u)
	if e != nil || i == 0 {
		log.Printf("插入数据到user失败 %v", e)
	}
	log.Printf("插入了 %d 条数据到user", i)
}

func ExcQuerySQL(sql string) (resultsSlice, err) {
	return XEngine.Query(sql)
}
