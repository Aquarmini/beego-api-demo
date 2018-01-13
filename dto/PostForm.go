package dto

type PostForm struct {
	Id   int         `form:"id"`
	Name interface{} `form:"name"`
}
