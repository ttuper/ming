package feedback

import (
	"github.com/gin-gonic/gin"
	"ming/internal/models"
	"net/http"
)

type submitRequest struct {
	Type int    `form:"type"` // 反馈类型：1=功能建议；2=使用问题；3=求资源；4=其他
	Desc string `form:"desc"` // 反馈内容
}

type submitResponse struct {
}

func (h *FeedbackHandler) SubmitFeedback(c *gin.Context) {
	var req submitRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	feedback := &models.Feedback{
		Type: req.Type,
		Desc: req.Desc,
	}

	err := h.FeedbackService.SubmitFeedback(feedback)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "submit ok!")

}
