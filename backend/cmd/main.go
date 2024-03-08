package main

import "backend/internal/app"

//	@title Revenue App
//	@version 1.0.0
//	@description API Server for Revenue App

//	@BasePath /api

func main() {

	app := app.New()
	//	Run
	app.Run()
}
