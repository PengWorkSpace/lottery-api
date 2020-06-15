package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"

	"lottery-api/internal/model"
	"lottery-api/utils"
)

func prizeList(c *gin.Context) {
	code := http.StatusOK
	data, err := svc.GetPrizeList(c)
	if err != nil {
		code = http.StatusNotFound
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}

//生成手机验证码存储到redis，用phone:guid 当作key，避免手机号验证码重复的情况
func generatePhoneCode(c *gin.Context) {
	//存储code到redis
	code := http.StatusOK
	phoneStr := c.Query("phone")
	fmt.Println(phoneStr)
	c.Render(code, render.JSON{Data: "23213sdfesfefef"})
}

func verifyPhoneCode(verifyCodeID string, verifyCode string) (isOk bool, message string) {
	if "23213sdfesfefef" == verifyCodeID {
		return true, "验证码校验正确"
	}
	return false, "验证码验证失败"
}

func drawPrize(c *gin.Context) {
	phoneStr := c.Query("phone")
	phone, _ := strconv.ParseInt(phoneStr, 10, 64)
	code := http.StatusOK
	drawRet, err := svc.DrawPrize(c, phone)
	if err != nil {
		code = http.StatusBadRequest
	}
	c.Render(code, render.JSON{Data: drawRet})

}

func drawRecords(c *gin.Context) {
	data := svc.ExportDrawRecords(c)
	c.Writer.Header().Set("Content-Type", "application/csv")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", data.Title))
	c.Render(http.StatusOK, utils.CSV{Title: data.Title, Content: data.Content})
}

func join(c *gin.Context) {
	code := http.StatusOK
	req := new(model.UserInfoReq)

	//校验验证码
	// verifyID := req.VerifyID
	// verifyCode := req.VerifyCode
	// verfify, message := verifyPhoneCode(verifyID, verifyCode)
	// if !verfify {
	// 	c.Render(http.StatusBadRequest, render.JSON{Data: message})
	// 	return
	// }

	if err := c.ShouldBind(&req); err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}

	verify, message := svc.VerifyUserPhone(c, req.Phone)
	if !verify {
		c.Render(http.StatusBadRequest, render.JSON{Data: message})
		return
	}

	data, err := svc.Join(c, req)
	if err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}

func userArticles(c *gin.Context) {
	code := http.StatusOK
	data, err := svc.UserArticles(c)
	if err != nil {
		c.Render(http.StatusBadRequest, render.JSON{Data: err})
		return
	}
	// TODO 错误响应规范化处理
	c.Render(code, render.JSON{Data: data})
}
