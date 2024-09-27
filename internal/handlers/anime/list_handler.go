package anime

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

type listRequest struct {
	Page     int    `form:"page"`      // 第几页
	PageSize int    `form:"page_size"` // 每页显示条数
	Keyword  string `form:"keyword"`   // 动漫片名
}

type listData struct {
	Id          int    `json:"id"`           // ID
	Title       string `json:"title"`        // 动漫片名
	Poster      string `json:"poster"`       // 动漫封面图
	Desc        string `json:"desc"`         // 简介
	Genres      string `json:"genres"`       // 动漫类型
	ReleaseDate string `json:"release_date"` // 上映日期
	Score       string `json:"score"`        // 评分
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// GetAnimeList converts the handler function to a middleware function for Gin.
func (h *AnimeHandler) GetAnimeList(c *gin.Context) {
	res := new(listResponse)
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	title := c.Query("title")

	page, _ := strconv.Atoi(pageStr)
	if page == 0 {
		page  = 1
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = 10
	}

	// 获取动漫列表
	animeList, err := h.AnimeService.GetAnimeList(page, pageSize, title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 获取动漫数量
	animeCount, err := h.AnimeService.GetAnimeCount(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get animes count failed"})
		return
	}

	res.Pagination.Total = animeCount
	res.Pagination.PerPageCount = pageSize
	res.Pagination.CurrentPage = page
	res.List = make([]listData, len(animeList))

	for k, v := range animeList {
		data := listData{
			Id:          cast.ToInt(v.ID),
			Title:       v.Title,
			Poster:      v.Poster,
			Desc:        v.Desc,
			Genres:      v.Genres,
			ReleaseDate: v.ReleaseDate,
			Score:       v.Score,
		}
		res.List[k] = data
	}

	c.JSON(http.StatusOK, res)
}
