package repo

import (
	"tianmiao-go/pkg/database"
)

// GetOne 获取单条数据
func GetOne(model interface{}, where [][]string, columns []string, order map[string]string) {
	//数据model结构体
	db := database.DB
	//查询字段
	db = BuildColumns(db, columns)
	//过滤条件
	db = BuildWhere(db, where)
	//排序
	db = BuildOrder(db, order)
	//获取一条记录，没有指定排序字段
	db = db.Take(model)
}

// GetById 通过主键获取单条数据
func GetById(model interface{}, id int, columns []string, order map[string]string) {
	//数据model结构体
	db := database.DB
	//查询字段
	db = BuildColumns(db, columns)
	//过滤条件
	db = db.Where("id =?", id)
	//排序
	db = BuildOrder(db, order)
	//获取一条记录，没有指定排序字段
	db = db.Take(model)
}

// GetRows 获取多条数据
func GetRows(model interface{}, where [][]string, columns []string, order map[string]string, limit int) {
	db := database.DB
	//查询字段
	db = BuildColumns(db, columns)
	//过滤条件
	db = BuildWhere(db, where)
	//排序
	db = BuildOrder(db, order)
	//limit
	db = BuildLimit(db, limit)
	//获取一条记录，没有指定排序字段
	db = db.Find(model)
}
