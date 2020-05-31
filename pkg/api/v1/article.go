package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_gin/models"
	"go_gin/pkg/e"
	"go_gin/pkg/response"
	"go_gin/pkg/setting"
	"go_gin/pkg/util"
)

// 获取全部articles
func GetArticles(c *gin.Context) {

	title := c.Query("title")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if title != "" {
		maps["title"] = title
	}
	state := c.Query("state")
	if state != "" {
		maps["state"] = com.StrTo(state).MustInt()
	}

	data["list"] = models.SelectArticleByPage(util.GetPage(c), setting.PageSize, maps)
	data["count"] = models.GetArticlesTotal(maps)

	response.ReturnJSON(c, e.SUCCESS, e.GetMsgFlag(e.SUCCESS), data)

}

func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0 {
		response.ReturnError(c,e.INVALID_PARAMS,e.GetMsgFlag(e.INVALID_PARAMS))
		return
	}

	article := models.SelectArticleById(id)

	response.ReturnJSON(c,e.SUCCESS,e.GetMsgFlag(e.SUCCESS),article)

}

// 新增article
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tagId")).MustInt()
	title := c.Query("title")
	content := c.Query("content")
	createdBy := c.Query("createdBy")
	state := com.StrTo(c.Query("state")).MustInt()

	errorMsg := CheckArticleParams(nil, title, content, state, nil, createdBy)
	if errorMsg != "" {
		response.ReturnError(c, e.INVALID_PARAMS, errorMsg)
		return
	}

	if models.ExistTagByName(title) {
		response.ReturnError(c, e.ERROR_NOT_EXIST_ARTICLE, e.GetMsgFlag(e.ERROR_NOT_EXIST_ARTICLE))
		return
	}

	models.AddArticle(&models.Article{
		TagId:     tagId,
		Title:     title,
		Content:   content,
		State:     state,
		CreatedBy: createdBy,
	})

	response.ReturnSuccess(c)

}

// 参数校验
func CheckArticleParams(id, title, content, state, modifiedBy, createdBy interface{}) string {
	validator := validation.Validation{}
	if id != nil {
		validator.Required(id, "id").Message("ID不能为空！")
	}
	if title != nil {
		validator.Required(title, "title").Message("Title不能为空！")
	}

	if content != nil {
		validator.Required(content, "content").Message("Content不能为空！")
	}
	if state != nil {
		validator.Required(state, "state").Message("state不能为空！")
		validator.Range(state, 0, 1, "state").Message("state只能为0或者1")
	}
	if modifiedBy != nil {
		validator.Required(modifiedBy, "modifiedBy不能为空")
	}
	if createdBy != nil {
		validator.Required(createdBy, "createdBy").Message("createdBy不能为空")
	}

	if validator.HasErrors() {
		return validator.Errors[0].Message
	}
	return ""

}
