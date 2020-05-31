package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 查询Tags
func SelectTags(pageNum int, pageSize int, maps interface{}) (tag []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tag)
	return
}

// 查询Tags count
func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// 是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID == 0 {
		return false
	} else {
		return true
	}

}

// 新增Tag
func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

// 编辑
func EditTag(id int, maps interface{}) {
	db.Model(&Tag{}).Where("id = ?", id).Update(maps)
}

// 删除
func DeleteTag(id int)  {
	db.Where("id = ?",id).Delete(&Tag{})
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}
