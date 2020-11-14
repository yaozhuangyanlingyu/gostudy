package options

import (
	"notify/pkg/logger"
)

var Opts options

// 全局配置
type options struct {
	Server GinServerConfig
	Log    logger.Config
}

// gin服务
type GinServerConfig struct {
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	HTTPAddr     string `mapstructure:"http_addr"`
	Mode         string
	Pprof        bool
}
