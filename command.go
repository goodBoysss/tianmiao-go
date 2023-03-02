package main

import (
	"flag"
	"fmt"
	"tianmiao-go/bootstrap"
	"tianmiao-go/pkg/config"
	"tianmiao-go/pkg/converter"
	"tianmiao-go/pkg/helpers"
)

func main() {

	var table string
	flag.StringVar(&table, "table", "", "--table")
	flag.Parse()

	//model := strings.ToUpper(table[:1]) + table[1:]
	model := helpers.UnderlineToCamel(table)
	createModel(table, model)

	////type Person struct {
	////	HelloWold       string
	////	LightWeightBaby string
	////}
	////var a = Person{HelloWold: "chenqionghe", LightWeightBaby: "muscle"}
	//res, _ := json.Marshal("aa_aa")
	//fmt.Printf("%s", res)

}

func createModel(table string, model string) {

	//应用初始化
	bootstrap.InitApp()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
		config.Get("database.mysql.username"),
		config.Get("database.mysql.password"),
		config.Get("database.mysql.host"),
		config.Get("database.mysql.port"),
		config.Get("database.mysql.database"),
		config.Get("database.mysql.charset"))

	// 初始化
	t2t := converter.NewTable2Struct()
	// 个性化配置
	t2t.Config(&converter.T2tConfig{
		// 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
		RmTagIfUcFirsted: false,
		// tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
		TagToLower: false,
		// 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
		UcFirstOnly: false,
		//// 每个struct放入单独的文件,默认false,放入同一个文件(暂未提供)
		//SeperatFile: false,
		//驼峰
		StructNameToHump: true,
	})
	// 开始迁移转换
	err := t2t.
		// 指定某个表,如果不指定,则默认全部表都迁移
		Table(table).
		// 生成的结构体保存路径
		SavePath("./app/Models/" + model + ".go").
		// 表前缀
		Prefix("ln_").
		// 是否添加json tag
		EnableJsonTag(true).
		// 生成struct的包名(默认为空的话, 则取名为: package model)
		PackageName("Models").
		// tag字段的key值,默认是orm
		TagKey("orm").
		// 是否添加结构体方法获取表名
		//RealNameMethod("TableName").
		DateToTime(true).

		// 数据库dsn,这里可以使用 t2t.DB() 代替,参数为 *sql.DB 对象
		Dsn(dsn).
		// 执行
		Run()

	fmt.Println(err)
}
