package dbgo

type Config struct {
	User         string
	Password     string
	Host         string
	Port         string
	DB           string
	MaxIdleConns int `mapstructure:"max_idle_conns"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	Log          bool
}
