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

	service.AuthRouter(router.Group("/auth"), &appCtx)

	router.Run(fmt.Sprintf(":%s", config.Port))
}
