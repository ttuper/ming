package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Anime struct {
	ID          int      `json:"id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Status      string    `json:"status"`
	Desc        string    `json:"desc"`
	Genres      string    `json:"genres"`
	Poster      string    `json:"poster"`
	ReleaseDate string    `json:"release_date"`
	Score       float64   `json:"score"`
}

// TableName sets the table name for the model.
func (Anime) TableName() string {
	return "animes"
}

func main() {
	//newLogger := logger.New(
	//	log.New(nil, "\r", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold:             time.Second, // Slow SQL threshold
	//		LogLevel:                  logger.Info, // Log level
	//		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
	//		Colorful:                  true,        // Disable color
	//	},
	//)

	// 创建日志配置
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的地方，这里是标准输出）
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level（日志级别）
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:      true,         // 彩色打印
		},
	)

	dsn := "root:123456@tcp(127.0.0.1:3306)/ming?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})

	var animes []*Anime
	query := db.Model(&Anime{})
	query.Find(&animes)
}


