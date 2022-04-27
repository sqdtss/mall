package models

// PrimaryKey 主键请求参数
type PrimaryKey struct {
	Id  uint   `form:"id"  json:"id"  binding:"required,gt=0"`
}

type WeiXinUser struct {
	Id  string   `form:"id"  json:"id"  binding:"required,gt=0"`
}

// Page 页面请求参数
type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"  binding:"required,gt=0"`
	PageSize int `form:"pageSize" json:"pageSize" binding:"required,gt=0"`
}
