package main

import (
	"log"
	"os"
	"os/signal"
	"scf-web-func-go-demo/app"
)

func init() {
	log.Println("App init...")
	a := app.GetApp()
	a.Init()
	a.Check()
}

func main() {
	app := app.GetApp()
	log.Println("App start...")
	app.Run()
	log.Println("App start handle async task...")
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Server force shutdown...")
	close(quit)
	os.Exit(1)
}
