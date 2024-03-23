package app

import (
	"backend/docs"
	"backend/internal/repository"
	"backend/internal/transport/http/handlers/middleware"
	"backend/internal/transport/http/handlers/revenue"
	"backend/internal/transport/http/handlers/users"
	"backend/internal/usecase"
	"backend/pkg/gorm/postgres"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"      // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@BasePath /api

type Application struct {
}

func New() *Application {
	app := &Application{}

	return app
}

const addr string = "user=postgres password=postgres host=172.30.144.79 port=5432 dbname=postgres sslmode=disable"

func (a *Application) Run() {
	err := postgres.RunMigrations(addr)
	if err != nil {
		log.Fatalf("failed to run migrations %v", err)
	}
	timeUTC := func() time.Time { return time.Now().UTC() }
	db, err := postgres.New(addr, postgres.TranslateError(true), postgres.NowFunc(timeUTC))
	if err != nil {
		log.Fatalf("postgres err %v", err)
	}
	repo := repository.NewRepository(db)

	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	v1 := r.Group("/api")
	usecase := usecase.New(repo)
	revenue.NewRevenueRoutes(v1, usecase)
	users.NewUsersRoutes(v1, usecase)

	docs.SwaggerInfo.BasePath = "/api"
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFile.Handler, "DISABLE_SWAGGER")
	r.GET("/swagger/*any", swaggerHandler)

	r.Run("localhost:8081")
}
