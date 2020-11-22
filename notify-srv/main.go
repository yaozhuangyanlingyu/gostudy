/*
 * @Author: yaoxf
 * @Date: 2020-09-13 14:24:39
 * @LastEditTime: 2020-11-22 22:33:54
 * @LastEditors: Please set LastEditors
 * @Description: 通知相关 - 服务接口
 * @FilePath: \notify-srv\main.go
 */
package main

import (
	"notify/api"
)

func main() {
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
