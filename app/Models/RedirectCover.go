package Models

import "time"

type RedirectCover struct {
	Id            int64     `orm:"id" json:"id"`                           // ID
	AppId         int64     `orm:"app_id" json:"app_id"`                   // 应用ID
	DomainMd5     string    `orm:"domain_md5" json:"domain_md5"`           // 域名md5
	ShortKey      string    `orm:"short_key" json:"short_key"`             // 短连接key
	CoverImageUrl string    `orm:"cover_image_url" json:"cover_image_url"` // 封面图片地址
	CreatedAt     time.Time `orm:"created_at" json:"created_at"`           // 创建时间
	UpdatedAt     time.Time `orm:"updated_at" json:"updated_at"`           // 更新时间
}
