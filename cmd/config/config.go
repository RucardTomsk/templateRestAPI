package config

import (
	"templateRestAPI/internal/common"
)

type Config struct {
	DB     common.DatabaseConfig
	Server common.ServerConfig
}
