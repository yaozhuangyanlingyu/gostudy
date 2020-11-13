package options

/**
 * 配置文件
 */
var Opts *options

type options struct {
	Server *GinServerConfig
}

type GinServerConfig struct {
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_time"`
	HTTPAddr     string `mapstructure:"http_addr"`
	Mode         string
	Pprof        bool
}
