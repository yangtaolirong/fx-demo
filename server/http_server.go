package server

import (
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/macaron.v1"
	"net/http"
	"server/pkg/config"
	"server/pkg/router"
)

type HttpServer struct {
	cfg * config.Config
	logger *zap.Logger
	mar * macaron.Macaron
}

func NewHttpServer(cfg * config.Config,logger *zap.Logger)*HttpServer  {
	return &HttpServer{
		cfg: cfg,
		logger: logger.Named("http_server"),
		mar:macaron.Classic() ,
	}
}
func (srv* HttpServer)Run()error  {
	router.Register(srv.mar.Router)
	addr:=fmt.Sprintf("0.0.0.0:%v",srv.cfg.HttpConfig.Port)
	srv.logger.Info("http run ",zap.String("addr",addr))
	return  http.ListenAndServe(addr, srv.mar)
}
