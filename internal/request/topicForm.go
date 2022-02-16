package request

// AddTopicForm 评论表单
type AddTopicForm struct {
	WorksId    int    `form:"worksId" valid:"Required;Min(1)"`
	Content    string `form:"content" valid:"Required"`
	RelationId int    `form:"relationId"`
}
