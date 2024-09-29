package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ming/internal/db"
	anime_handler "ming/internal/handlers/anime"
	anime_service "ming/internal/service/anime"
	"ming/pkg/config"
	"ming/pkg/logger"
	"net/http"
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

	// 设置TLS配置
	config := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256, // 适用于2021年1月及之后的实例
			// tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256, // 可选，适用于2021年1月之前的实例
			// tls.TLS_RSA_WITH_AES_128_CBC_SHA256,      // 可选，适用于2021年1月之前的实例
		},
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
	}

	// 添加路由规则
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动HTTPS服务器
	server := &http.Server{
		Addr:      ":443",
		Handler:   r,
		TLSConfig: config,
	}

	log.Println("Starting server...")
	if err := server.ListenAndServeTLS("/root/ssl/www.mingcy.fun.pem", "/root/ssl/www.mingcy.fun.key"); err != nil {
		log.Fatalf("ListenAndServeTLS error: %v", err)
	}
}
