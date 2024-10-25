package vaccine

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type listRequest struct {
	Keyword string `form:"keyword"` // 疫苗名称
}

type listData struct {
	Id          int    `json:"id"`          // ID
	Name        string `json:"name"`        // 疫苗名称
	FullName    string `json:"full_name"`   // 疫苗全称
	Category    string `json:"category"`    // 疫苗类别
	Doses       string `json:"doses"`       // 接种时间点
	TotalDoses  int    `json:"total_doses"` // 总接种次数
	Description string `json:"description"` // 简介
	IsFree      int    `json:"is_free"`
}

type listResponse struct {
	List []listData `json:"list"`
}

// GetVaccineList 获取所有疫苗列表
func (h *VaccineHandler) GetVaccineList(c *gin.Context) {
	res := new(listResponse)
	keyword := c.Query("keyword")

	// 获取疫苗列表
	vaccineList, err := h.VaccineService.GetVaccineList(keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res.List = make([]listData, len(vaccineList))

	for k, v := range vaccineList {
		// 将逗号分隔的字符串doses转换为切片
		doses := strings.ReplaceAll(v.Doses, ",", "、")
		totalDoses := len(strings.Split(v.Doses, ","))
		data := listData{
			Id:          cast.ToInt(v.ID),
			Name:        v.Name,
			FullName:    v.FullName,
			Category:    v.Category,
			Doses:       doses,
			TotalDoses:  totalDoses,
			Description: v.Description,
		}
		res.List[k] = data
	}

	c.JSON(http.StatusOK, res)
}
