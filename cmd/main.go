package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ming/internal/db"
	anime_handler "ming/internal/handlers/anime"
	anime_service "ming/internal/service/anime"
	"ming/pkg/config"
	"ming/pkg/logger"
)

func main() {
	var (
		configFile string
	)
	flag.StringVar(&configFile, "c", "", "Configuration file path.")
	flag.StringVar(&configFile, "config", "", "Configuration file path.")
	flag.Parse()

	// 初始化配置
	config.InitConfig(configFile)
	cfg := config.GetConfig()

	// 初始化数据库连接
	if err := db.ConnectMySQL(cfg); err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to MySQL: %v", err))
		return
	}

	//if err := db.ConnectRedis(cfg); err != nil {
	//	logger.Error(fmt.Sprintf("Failed to connect to Redis: %v", err))
	//	return
	//}

	// 初始化路由
	r := gin.Default()
	animeService := anime_service.NewAnimeService()
	animeHandler := anime_handler.NewAnimeHandler(animeService)

	anime := r.Group("/anime")
	{
		anime.GET("/list", animeHandler.GetAnimeList)
		anime.GET("/detail/:anime_id", animeHandler.GetAnimeByID)
	}

	fmt.Println("Listening on :443...")
	logger.Info("Server started.")
	if err := r.RunTLS(":443", "/root/ssl/www.mingcy.fun.pem", "/root/ssl/www.mingcy.fun.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
