package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin/pkg/api/v1"
	"go_gin/pkg/e"
	"go_gin/pkg/setting"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsgFlag(e.SUCCESS),
		})
	})

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tag", v1.AddTag)
		apiV1.PUT("/tag/:id", v1.EditTag)
		apiV1.DELETE("/tag/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
	}

	return r
}
