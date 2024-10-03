package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
	//r := gin.Default()
	//animeService := anime_service.NewAnimeService()
	//animeHandler := anime_handler.NewAnimeHandler(animeService)
	//
	//anime := r.Group("/anime")
	//{
	//	anime.GET("/list", animeHandler.GetAnimeList)
	//	anime.GET("/detail/:anime_id", animeHandler.GetAnimeByID)
	//}

	// 1. 加载证书文件
	certFile := "/etc/ssl/certs/aliyun-root-ca.crt"
	certBytes, err := ioutil.ReadFile(certFile)
	if err != nil {
		log.Fatalf("Failed to read certificate file: %v", err)
	}

	// 2. 解析证书
	rootCAs := x509.NewCertPool()
	if !rootCAs.AppendCertsFromPEM(certBytes) {
		log.Fatalf("Failed to append cert to pool")
	}

	// 3. 创建 TLS 配置
	tlsConfig := &tls.Config{
		RootCAs: rootCAs,
		MinVersion:  tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		},
	}

	//// 4. 初始化 Gin 框架
	//router := gin.Default()
	//
	//// 5. 添加路由
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "Hello, world!",
	//	})
	//})

	r := gin.Default()
	animeService := anime_service.NewAnimeService()
	animeHandler := anime_handler.NewAnimeHandler(animeService)

	anime := r.Group("/anime")
	{
		anime.GET("/list", animeHandler.GetAnimeList)
		anime.GET("/detail/:anime_id", animeHandler.GetAnimeByID)
	}

	// 6. 创建 HTTP 服务器实例
	server := &http.Server{
		Addr:      ":443", // 你可以更改端口号
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	// 7. 启动 HTTPS 服务
	log.Println("Starting HTTPS server...")
	err = server.ListenAndServeTLS("", "")
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}


	//fmt.Println("Listening on :443...")
	//logger.Info("Server started.")
	//if err := r.RunTLS(":443", "/ssl/cert.pem", "/ssl/cert.key"); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}
}
