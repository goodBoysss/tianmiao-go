package config

import (
	"tianmiao-go/pkg/config"
)

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认数据库
			"connection": config.Env("DB_CONNECTION", "mysql"),
			//数据库表名称前缀
			"table_prefix": config.Env("DB_TABLE_PREFIX", "ln_"),

			"mysql": map[string]interface{}{

				// 数据库连接信息
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "tianmiao"),
				"username": config.Env("DB_USERNAME", "root"),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  config.Env("DB_CHARSET", "utf8mb4"),

				// 连接池配置
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
		}
	})
}
