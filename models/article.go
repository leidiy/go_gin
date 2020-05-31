package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model
	TagId int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"describe"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// 分页查询Articles
func SelectArticleByPage(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}
func GetArticlesTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func SelectArticleById(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	// 将tag 关联到article上
	db.Model(&article).Related(&article.Tag)
	return
}

func ExistArticleByTitle(title string) bool {
	tag := Tag{}
	db.Where("title = ?", title).First(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

// 新增Article
func AddArticle(article *Article) {
	db.Create(article)
	return
}

func (a *Article) BeforeCreate(scope *gorm.Scope) {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}

func (a *Article) BeforeUpdate(scope *gorm.Scope) {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
}
