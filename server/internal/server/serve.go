package server

import (
	"fmt"
	"net/http"
	model "server/internal/config"
	"server/internal/route"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 单个实现无需定义接口

// var _ Server = (*Serve)(nil)

// type Server interface {
// 	InitServer(*gorm.DB, *zap.Logger, *model.Config)
// 	RegisterServer()
// 	RunServer(*gorm.DB, *zap.Logger, *model.Config) error
// }

type Serve struct {
	Mysql  *gorm.DB
	Logger *zap.Logger
	Engine *gin.Engine
	Config model.System
	// routerGroup *gin.RouterGroup
}

func New(db *gorm.DB, logger *zap.Logger) *Serve {
	return &Serve{
		Mysql:  db,
		Logger: logger,
	}
}

func (s *Serve) InitServer(db *gorm.DB, logger *zap.Logger, config *model.Config) {
	e := gin.Default()
	// gin中间件加载
	e.Use()
	gin.SetMode(config.System.Mode)
	// gin.SetMode(gin.ReleaseMode)
	// gin.ReleaseMode
	s.Engine = e
	s.Config = config.System
}

func (s *Serve) RegisterServer() {
	routerGroup := s.Engine.Group("")
	{
		route.InitRouter(routerGroup, s.Engine)
	}
}

func (s *Serve) RunServer(db *gorm.DB, logger *zap.Logger, config *model.Config) error {
	s = New(db, logger)
	s.InitServer(db, logger, config)
	s.RegisterServer()
	h := http.Server{
		Addr:         ":" + s.Config.Port,
		Handler:      s.Engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	logger.Info(fmt.Sprintf("Server listen port:%s", s.Config.Port))
	return h.ListenAndServe()
	// return s.Engine.Run(":8080")
}
