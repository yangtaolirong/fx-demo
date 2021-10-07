package server

import (
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	group errgroup.Group //errgroup，参考我的文章，专门讲这个原理
	app *fx.App //fx 实例
	provides []interface{}
	invokes  []interface{}
	supplys   []interface{}
	httSrv *HttpServer //该http server 可以换成fibber gin 之类的
}

func NewServer(
	)*Server  {
	return &Server{

	}
}
func(srv*Server) Run()  {
	srv.app=fx.New(
		fx.Provide(srv.provides...),
		fx.Invoke(srv.invokes...),
		fx.Supply(srv.supplys...),
		fx.Provide(NewHttpServer),//注入http server
		fx.Supply(srv),
		fx.Populate(&srv.httSrv), //给srv 实例赋值
		fx.NopLogger,//禁用fx 默认logger
		)
	srv.group.Go(srv.httSrv.Run) //启动http 服务器
	err:=srv.group.Wait() //等待子协程退出
	if err!=nil{
		panic(err)
	}
}
func(srv*Server)Provide(ctr  ...interface{}){
	srv.provides= append(srv.provides, ctr...)
}
func(srv*Server)Invoke(invokes  ...interface{}){
   srv.invokes=append(srv.invokes,invokes...)
}
func(srv*Server)Supply(objs ...interface{}){
   srv.supplys=append(srv.supplys,objs...)
}
