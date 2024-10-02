package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
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

	fmt.Println("Listening on :443...")
	logger.Info("Server started.")
	//if err := r.RunTLS(":443", "/root/ssl/www.mingcy.fun.pem", "/root/ssl/www.mingcy.fun.key"); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}

	// 加载TLS配置
	tlsCfg := &tls.Config{
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		},
		PreferServerCipherSuites: true,
		InsecureSkipVerify:       false,
		MinVersion:               tls.VersionTLS12,
		MaxVersion:               tls.VersionTLS13,
	}

	// 使用TLS配置创建一个TLS Listener
	server := &http.Server{
		Addr:      ":443", // 监听端口
		Handler:   r,
		TLSConfig: tlsCfg,
	}

	// 启动TLS服务器
	err := server.ListenAndServeTLS("/root/ssl/www.mingcy.fun.pem", "/root/ssl/www.mingcy.fun.key")
	if err != nil {
		panic(err)
	}
}
