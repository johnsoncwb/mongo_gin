package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/mongo_gin/internal/server/handlers"
	"log"
	"os"
)

type Application struct {
	router *gin.Engine
}

func Init() *Application {
	router := gin.Default()

	router.POST("/user", handlers.SaveOnDatabase)
	router.GET("/user/:id", handlers.GetById)
	router.GET("/user/condition/:maried", handlers.GetAllByCondition)

	return &Application{router: router}
}

func (a *Application) Start() {
	log.Fatal(a.router.Run(fmt.Sprintf("localhost:%s", os.Getenv("PORT"))))
}
