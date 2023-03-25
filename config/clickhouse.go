package config

import (
	"base-api/pkg/config"
)

func init() {

	config.Add("clickhouse", func() map[string]interface{} {
		return map[string]interface{}{
			"ip":       config.Env("CLICKHOUSE_IP", "127.0.0.1"),
			"port":     config.Env("CLICKHOUSE_PORT", "9000"),
			"user":     config.Env("CLICKHOUSE_USER", "default"),
			"password": config.Env("CLICKHOUSE_PASSWORD", "guest"),
			"database": config.Env("CLICKHOUSE_DATABASE", "guest"),
		}
	})
}
