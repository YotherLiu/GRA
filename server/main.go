package main

import (
	"fmt"
	"server/internal/pkg/db"
	"server/internal/pkg/viper"
	"server/internal/pkg/zap"
	"server/internal/server"
)

func main() {
	// 配置读取
	var v viper.Viper
	config, err := v.InitViper()
	if err != nil {
		panic(fmt.Sprintf("Configuration error：%s", err.Error()))
	}
	// 日志配置
	z := zap.Logger{}
	logger := z.InitLogger(config)
	// 数据库连接
	g := db.Mysql{}
	db, err := g.InitMysql(config)
	if err != nil {
		panic(fmt.Sprintf("Database initialization failed：%s", err.Error()))
	}
	if err := g.InitMigrate(db); err != nil {
		if d, err := db.DB(); err != nil {
			defer d.Close()
		}
	} else {
		logger.Info("Database migration completed.")
	}
	// gin服务注册
	s := server.Serve{}
	if err := s.RunServer(db, logger, config); err != nil {
		panic(fmt.Sprintf("Server start failed：%s", err.Error()))
	}
}
