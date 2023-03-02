package Models

import "time"

type RedirectUrl struct {
	Id          int64     `orm:"id" json:"id"`                       // ID
	AppId       int64     `orm:"app_id" json:"app_id"`               // 应用ID
	DomainMd5   string    `orm:"domain_md5" json:"domain_md5"`       // 域名md5
	ShortKey    string    `orm:"short_key" json:"short_key"`         // 短连接key
	OriginUrl   string    `orm:"origin_url" json:"origin_url"`       // 原始链接url
	IsShowCover int64     `orm:"is_show_cover" json:"is_show_cover"` // 是否展示封面图（微信、QQ）：1-展示；0-不展示；
	CreatedAt   time.Time `orm:"created_at" json:"created_at"`       // 创建时间
	UpdatedAt   time.Time `orm:"updated_at" json:"updated_at"`       // 更新时间
}
