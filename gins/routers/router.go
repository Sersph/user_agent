package routers

import (
	"bytes"
	"gins/controller"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	// 确保发生错误时不会崩溃
	r.Use(gin.Recovery())

	r.Use(Cors())

	r.POST("/v1/data", controller.ProcessData)
	r.POST("/v1/data/edit", controller.Edit)
	r.POST("/v1/setCookies", controller.SetCookies)
	r.GET("/v1/getCookies", controller.GetCookies)
	r.GET("/v1/data/list", controller.List)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// 自定义的ResponseWriter
type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写Write方法，捕获响应内容
func (w responseWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

// 获取请求Body内容
func getRequestBody(c *gin.Context) string {
	body, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return string(body)
}
