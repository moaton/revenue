package main

import "backend/internal/app"

// @title Revenue App
// @version 1.0.0
// @description API Server for Revenue App

// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	app := app.New()
	//	Run
	app.Run()
}
