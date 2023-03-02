package database

import (
	"encoding/json"
	"gorm.io/gorm"
	"strings"
	"tianmiao-go/pkg/helpers"
)

type Repo struct {
	model interface{}
}

func (repo *Repo) Source(model interface{}) *Repo {
	repo.model = model
	return repo
}

// GetOne 获取单条数据
func (repo *Repo) GetOne(where [][]string, columns []string, order map[string]string) map[string]interface{} {
	//数据model结构体
	model := repo.model
	db := DB
	//查询字段
	db = BuildColumns(db, columns)
	//过滤条件
	db = BuildWhere(db, where)
	//排序
	db = BuildOrder(db, order)
	//获取一条记录，没有指定排序字段
	db = db.Take(&model)

	//处理返回格式，返回map
	return dealRes(model, columns)
}

// GetOne 获取多条数据
func (repo *Repo) GetRows(where [][]string, columns []string, order map[string]string) map[string]interface{} {
	//数据model结构体
	model := repo.model
	db := DB
	//查询字段
	db = BuildColumns(db, columns)
	//过滤条件
	db = BuildWhere(db, where)
	//排序
	db = BuildOrder(db, order)
	//获取一条记录，没有指定排序字段
	db = db.Take(&model)

	return helpers.StructToMap(model)
}

// BuildColumns 构建查询字段
func BuildColumns(db *gorm.DB, columns []string) *gorm.DB {
	//查询字段
	if columns != nil {
		db = db.Select(columns)
	}
	return db
}

// BuildWhere 过滤条件
func BuildWhere(db *gorm.DB, where [][]string) *gorm.DB {
	for _, list := range where {

		column := list[0]
		operator := list[1]
		value := list[2]

		if column != "" && operator != "" {
			operator = strings.ToUpper(operator)
			if operator == "BETWEEN" {
				value2 := list[3]
				db.Where(column+" "+operator+" ? AND ?", value, value2)
			} else {
				db.Where(column+" "+operator+" ?", value)
			}

		}
	}
	return db
}

// BuildOrder 构建排序
func BuildOrder(db *gorm.DB, order map[string]string) *gorm.DB {
	orderStr := ""
	//排序
	if order != nil {
		for colums, sort := range order {
			if colums != "" && sort != "" {
				if orderStr == "" {
					orderStr = colums + " " + sort
				} else {
					orderStr += "," + colums + " " + sort
				}
			}
		}
		db = db.Order(orderStr)
	}
	return db
}

func dealRes(model interface{}, columns []string) map[string]interface{} {

	//先转json，再转回对象
	dat1, _ := json.Marshal(model)
	var dat2 map[string]interface{}
	json.Unmarshal(dat1, &dat2)

	//结构体转集合
	res := helpers.StructToMap(dat2)

	//剔除select外的字段
	var newRes = make(map[string]interface{})
	for _, column := range columns {
		if column == "*" {
			newRes = res
			return newRes
		} else {
			newRes[column] = res[column]
		}
	}
	return newRes
}
