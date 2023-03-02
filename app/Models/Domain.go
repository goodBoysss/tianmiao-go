package Models

import "time"

type Domain struct {
	Id          int64     `orm:"id" json:"id"`                     // ID
	AppId       int64     `orm:"app_id" json:"app_id"`             // 应用ID
	Domain      string    `orm:"domain" json:"domain"`             // 域名（不带http及/）
	DomainMd5   string    `orm:"domain_md5" json:"domain_md5"`     // 域名md5
	IsPublished int64     `orm:"is_published" json:"is_published"` // 是否开启：1-开启；0-未开启；
	CreatedAt   time.Time `orm:"created_at" json:"created_at"`     // 创建时间
	UpdatedAt   time.Time `orm:"updated_at" json:"updated_at"`     // 更新时间
}
