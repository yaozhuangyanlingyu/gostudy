package main

import (
	"notify/api"
	"notify/pkg/config"

	"notify/pkg/logger"
	"notify/pkg/options"
)

func init() {
	// 初始化配置
	v, err := config.SetupConfig("./configs")
	if err != nil {
		panic(err)
	}
	config.Decode(v.AllSettings(), &options.Opts)

	// 初始化日志库
	logger.SetupLogger(&options.Opts.Log)
}

func main() {
	logger.Info("starting service...")
	api.RunGin(&options.Opts.Server)
}
