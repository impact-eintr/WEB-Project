package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint   `json:"state"`
}

func (this *Tag) TableName() string {
	return "blog_tag"
}
