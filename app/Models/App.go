package Models

import "time"

type App struct {
	Id          int64     `orm:"id" json:"id"`                     // ID
	Name        string    `orm:"name" json:"name"`                 // 名称
	Alias       string    `orm:"alias" json:"alias"`               // 应用别名
	AppId       string    `orm:"app_id" json:"app_id"`             // 应用ID
	AppSecret   string    `orm:"app_secret" json:"app_secret"`     // 应用秘钥
	IsPublished int64     `orm:"is_published" json:"is_published"` // 是否开启：0-未开启；1-已开启；
	IsDeleted   int64     `orm:"is_deleted" json:"is_deleted"`     // 是否删除：0-未删除；1-已删除
	CreatedAt   time.Time `orm:"created_at" json:"created_at"`     // 创建时间
	UpdatedAt   time.Time `orm:"updated_at" json:"updated_at"`     // 更新时间
}
