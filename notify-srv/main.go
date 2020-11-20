package main

import (
	"notify/api"
)

func main() {
	// 初始化路由
	e := api.Engine{}
	e.InitRouter()
	e.RunEngine()

	// 运行框架
	/*
		serverConf := options.Opts.Server
		server := &http.Server{
			Addr:           serverConf.HTTPAddr,
			Handler:        router,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		go func() {
			if err := server.ListenAndServe(); err != nil && err.Error() != "http: Server closed" {
				logger.Error(fmt.Sprintf("failed to start server，error: %s", err.Error()))
			}
		}()*/
}
