package vaccine

import (
	"github.com/gin-gonic/gin"
	"strings"

	"net/http"
	"strconv"
)

type detailRequest struct {
	Id int32 `uri:"id"`
}

type detailResponse struct {
	Id          int32    `json:"id"`          // ID
	Name        string   `json:"name"`        // 疫苗名称
	FullName    string   `json:"full_name"`   // 疫苗全称
	Category    string   `json:"category"`    // 疫苗类别
	Doses       []string `json:"doses"`       // 接种时间点
	TotalDoses  int      `json:"total_doses"` // 接种针数
	Description string   `json:"description"` // 简介
	Detail      []struct {
		Title   string `json:"title"`   // 副标题
		Content string `json:"content"` // 内容
	} `json:"detail"` // 详细说明
	IsFree int32 `json:"is_free"`
}

// GetVaccineByID 获取疫苗的详细信息
func (h *VaccineHandler) GetVaccineByID(c *gin.Context) {
	res := new(detailResponse)
	idStr := c.Param("vaccine_id")
	id, _ := strconv.Atoi(idStr)

	vaccine, vaccineDetails, err := h.VaccineService.GetVaccineByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if vaccine.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vaccine not found"})
		return
	}

	doses := strings.Split(vaccine.Doses, ",")
	totalDoses := len(strings.Split(vaccine.Doses, ","))
	res.Id = int32(vaccine.ID)
	res.Name = vaccine.Name
	res.FullName = vaccine.FullName
	res.Category = vaccine.Category
	res.Doses = doses
	res.TotalDoses = totalDoses
	res.Description = vaccine.Description
	res.IsFree = int32(vaccine.IsFree)

	// 构建 Detail 字段
	for _, d := range vaccineDetails {
		res.Detail = append(res.Detail, struct {
			Title   string `json:"title"`
			Content string `json:"content"`
		}{
			Title:   d.Title,
			Content: d.Content,
		})
	}

	c.JSON(http.StatusOK, res)
}
