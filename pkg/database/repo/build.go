package repo

import (
	"gorm.io/gorm"
	"strings"
)

// BuildColumns 构建查询字段
func BuildColumns(db *gorm.DB, columns []string) *gorm.DB {
	//查询字段
	if columns == nil {
		columns = []string{"*"}
	}
	db = db.Select(columns)
	return db
}

// BuildWhere 过滤条件
func BuildWhere(db *gorm.DB, where [][]string) *gorm.DB {
	if where != nil {
		for _, list := range where {

			column := list[0]
			operator := list[1]
			value := list[2]

			if column != "" && operator != "" {
				operator = strings.ToUpper(operator)
				if operator == "BETWEEN" {
					value2 := list[3]
					db.Where(column+" BETWEEN ? AND ?", value, value2)
				} else if operator == "IN" {
					arr := list[2:]
					db.Where(column+" IN ?", arr)
				} else {
					db.Where(column+" "+operator+" ?", value)
				}
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
	} else {
		return db
	}
	return db
}

// BuildLimit 限制条件
func BuildLimit(db *gorm.DB, limit int) *gorm.DB {
	//排序
	db = db.Limit(limit)
	return db
}

//
//// 处理结构体
//func dealStruct(model interface{}, columns []string) map[string]interface{} {
//
//	//先转json，再转回对象
//	//dat1, _ := json.Marshal(model)
//	//var dat2 map[string]interface{}
//	//json.Unmarshal(dat1, &dat2)
//
//	//结构体转集合
//	//res := helpers.StructToMap(model)
//	//fmt.Println(res)
//	//fmt.Println(res)
//	////剔除select外的字段
//	//var newRes = make(map[string]interface{})
//	//for _, column := range columns {
//	//	if column == "*" {
//	//		newRes = res
//	//		return newRes
//	//	} else {
//	//		newRes[column] = res[column]
//	//	}
//	//}
//	//return newRes
//
//	var newRes = make(map[string]interface{})
//
//	modelVal := reflect.ValueOf(model)
//	for _, column := range columns {
//		//下划线转驼峰
//		columnCamel := helpers.UnderlineToCamel(column)
//		newRes[column] = modelVal.FieldByName(columnCamel)
//		fmt.Println(modelVal.Type())
//	}
//	return newRes
//}
//
//// 处理结构体
//func dealSliceStruct(models interface{}, columns []string) []map[string]interface{} {
//
//	modelVal := reflect.ValueOf(models)
//	modelLen := modelVal.Len()
//
//	var newRes = make([]map[string]interface{}, modelLen)
//	for i := 0; i < modelLen; i++ {
//		newRes[i] = dealStruct(modelVal.Index(i).Interface(), columns)
//	}
//
//	return newRes
//}
