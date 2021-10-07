package router

import "gopkg.in/macaron.v1"

func Register(router * macaron.Router)  {
	router.Get("/hello", func(ctx *macaron.Context) {
		ctx.Write([]byte("hello"))
	})
}
