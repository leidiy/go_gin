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

// 获取全部tags
func GetTags(c *gin.Context) {

	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	state := c.Query("state")
	if state != "" {
		maps["state"] = com.StrTo(state).MustInt()
	}

	data["list"] = models.SelectTags(util.GetPage(c), setting.PageSize, maps)
	data["count"] = models.GetTagsTotal(maps)

	response.ReturnJSON(c, e.SUCCESS, e.GetMsgFlag(e.SUCCESS), data)

}

// 新增Tag
func AddTag(c *gin.Context) {
	validate := validation.Validation{}
	name := c.PostForm("name")
	validate.Required(name, "name").Message("名称不能为空！")
	validate.MaxSize(name, 100, "name").Message("名称长度不能超过100位！")
	state := com.StrTo(c.PostForm("state")).MustInt()
	validate.Required(state, "state").Message("状态值不能为空！")
	validate.Range(state, 0, 1, "name").Message("状态值只能是0或者1！")
	createdBy := c.PostForm("createdBy")
	validate.Required(createdBy, "createdBy").Message("创建人不能为空！")
	var msg = ""
	var code = 0
	if validate.HasErrors() {
		code = e.INVALID_PARAMS
		msg = validate.Errors[0].Message
		response.ReturnError(c, code, msg)
		return
	}
	if models.ExistTagByName(name) {
		code = e.ERROR_EXIST_TAG
		msg = e.GetMsgFlag(code)
		response.ReturnError(c, code, msg)
		return
	}

	models.AddTag(name, state, createdBy)

	response.ReturnSuccess(c)

}

// 编辑Tag
func EditTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.Query("state")).MustInt()
	modifiedBy := c.Query("modifiedBy")
	id := com.StrTo(c.Param("id")).MustInt()

	errorMsg := CheckParams(id, name, state, modifiedBy, nil)
	if errorMsg != "" {
		response.ReturnError(c, e.INVALID_PARAMS, errorMsg)
		return
	}

	maps := make(map[string]interface{})
	maps["name"] = name
	maps["state"] = state
	maps["modifiedBy"] = modifiedBy
	models.EditTag(id, maps)

	response.ReturnSuccess(c)
}

// 参数校验
func CheckParams(id, name, state, modifiedBy, createdBy interface{}) string {
	validator := validation.Validation{}
	if id != nil {
		validator.Required(id, "id").Message("ID不能为空！")
	}
	if name != nil {
		validator.Required(name, "name").Message("Name不能为空！")
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

// 删除Tag
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if id == 0{
		response.ReturnError(c,e.INVALID_PARAMS,e.GetMsgFlag(e.INVALID_PARAMS))
		return
	}

	models.DeleteTag(id)
	response.ReturnSuccess(c)
}
