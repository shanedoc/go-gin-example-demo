package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model
	TagID int `json:"tag_id" gorm:"index"` //索引字段
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	//当前时间戳
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleById(id int) bool {
	var article Article
	db.Select("id").Where("id=?", id).First(&article)
	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

//列表
func GetArticles(pageNum int, PageSize int, maps interface{}) (articles []Article) {
	//预查询：查询出后会根据结构进行数据映射，避免循环查询
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(PageSize).Find(&articles)
	return
}

//详情
func GetArticle(id int) (article Article) {
	db.Where("id=?", id).First(&article)
	//关联查询
	db.Model(&article).Related(&article.Tag)
	return
}

//编辑
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id=?", id).Update(data)
	return true
}

//新增
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

//删除
func DeleteArticle(id int) bool {
	db.Where("id=?", id).Delete(Article{})
	return true
}
