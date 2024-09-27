package anime

import (
	"github.com/gin-gonic/gin"
	"strconv"

	"net/http"
)

type detailRequest struct {
	Id int32 `uri:"id"`
}

type ResourceList struct {
	Id      int32  `json:"id"`
	AnimeId int32  `json:"anime_id"`
	Url     string `json:"url"`
	Code    string `json:"code"`
	Type    string `json:"type"`
	UserId  int32  `json:"user_id"`
}

type detailResponse struct {
	Id          int32          `json:"id"`           // ID
	Title       string         `json:"title"`        // 动漫片名
	Poster      string         `json:"poster"`       // 动漫封面图
	Desc        string         `json:"desc"`         // 简介
	Genres      string         `json:"genres"`       // 动漫类型
	ReleaseDate string         `json:"release_date"` // 上映日期
	Score       string         `json:"score"`        // 评分
	Resources   []ResourceList `json:"resources"`
}

// GetAnimeByID converts the handler function to a middleware function for Gin.
func (h *AnimeHandler) GetAnimeByID(c *gin.Context) {
	res := new(detailResponse)
	idStr := c.Param("anime_id")
	id, _ := strconv.Atoi(idStr)

	// 获取 anime 数据
	anime, resources, err := h.AnimeService.GetAnimeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if anime.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Anime not found"})
		return
	}

	//组织资源数据
	resourceListData := make([]ResourceList, 0, len(resources))
	for _, v := range resources {
		// 资源类型：1=百度云；2=阿里云；3=夸客云；4=迅雷云；5=移动云；6=其他
		var resourceTypeName string
		switch v.Type {
		case 1:
			resourceTypeName = "百度云"
		case 2:
			resourceTypeName = "夸客云"
		case 3:
			resourceTypeName = "阿里云"
		case 4:
			resourceTypeName = "迅雷云"
		case 5:
			resourceTypeName = "移动云"
		default:
			resourceTypeName = "其他"
		}
		resourceListData = append(resourceListData, ResourceList{
			Id: int32(v.ID),
			AnimeId: v.AnimeID,
			Url:     v.Url,
			Code:    v.Code,
			Type:    resourceTypeName,
		})
	}

	res.Id = int32(id)
	res.Title = anime.Title
	res.Poster = anime.Poster
	res.Desc = anime.Desc
	res.Genres = anime.Genres
	res.ReleaseDate = anime.ReleaseDate
	res.Score = anime.Score
	res.Resources = resourceListData

	c.JSON(http.StatusOK, res)
}
