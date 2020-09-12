package options

import (
	"notify/tools/dbgo"
)

var Opts *options

type options struct {
	Server   *GinServerConfig
	AplumDb  *dbgo.Config    `mapstructure:"aplum_db"`
	ShortUrl *ShortUrlConfig `mapstructure:"short_url"`
}

type GinServerConfig struct {
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
	HTTPAddr     string `mapstructure:"http_addr"`
	Mode         string
	Pprof        bool
}

type ShortUrlConfig struct {
	BaseUrl string `mapstructure:"base_url"`
}
