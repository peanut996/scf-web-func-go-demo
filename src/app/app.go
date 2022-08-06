package app

import (
	"log"
	"scf-web-func-go-demo/api"
	"scf-web-func-go-demo/constant"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	once sync.Once
	app  *App
)

type App struct {
	srv *api.Server
	env *constant.Env
}

func GetApp() *App {
	once.Do(func() {
		a := &App{}
		app = a
	})
	return app
}

func (a *App) Check() {
}

func (a *App) Init() {
	log.Println("App init...")
	e := &constant.Env{}
	e.Init()
	a.env = e

	if "prod" == e.Stage {
		gin.SetMode(gin.ReleaseMode)
	}

	srv := api.NewServer()
	srv.Init()
	a.srv = srv

}

func (a *App) Run() {
	log.Println("App run...")
	a.srv.Run()
}
