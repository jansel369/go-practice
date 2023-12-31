package main

import (
	"fmt"
	"runtime"

	pgorm "server/serverless"
	service "server/services"
	utils "server/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cpuCount := runtime.NumCPU()

	fmt.Println("Number of cpu: ", cpuCount)

	config := utils.ReadConfig()
	orm := pgorm.InitOrm(&config.PgConfig)
	appCtx := utils.AppCtx{
		ORM: orm,
	}

	router := gin.Default()

	if config.ENV == "development" {
		router.SetTrustedProxies([]string{"localhost", "0.0.0.0"})
	}

	apiGroup := router.Group("/api/v1")

	service.AuthRouter(apiGroup.Group("/auth"), &appCtx)

	router.Run(fmt.Sprintf(":%s", config.Port))
}
