package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"ming/internal/db"
	"ming/internal/models"
	"ming/pkg/config"
	"ming/pkg/logger"
	"os"
	"strconv"
)

func main () {
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

	file, err := os.Open("zy.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// title,poster,desc,genres,score,release_date,type,url,code,user_i
	// 3. 插入数据到数据库
	// 3. 插入数据到数据库
	for _, record := range records {
		title := record[0]
		poster := record[1]
		desc := record[2]
		genres := record[3]
		score := record[4]
		releaseDate := record[5]
		url := record[6]
		code := record[7]
		typeResource := record[8] // 资源类型

		// 创建 Anime 记录
		anime := models.Anime{
			Title:         title,
			Poster:        poster,
			Desc:          desc,
			Genres:        genres,
			Score:         score,
			ReleaseDate:   releaseDate,
		}

		db.DB.Create(&anime)

		// 创建 Resource 记录
		typeR, _ := strconv.Atoi(typeResource)
		resource := models.Resource{
			AnimeID: int32(anime.ID),
			Type: int8(typeR),
			Url:     url,
			Code:    code,
		}

		db.DB.Create(&resource)

		fmt.Println("Record inserted successfully:", record)
	}

}
