package model

type Post struct {
	Body     string `json:"body"`
	CreateBy string `json:"created_by"`
	CreateAt int64  `json:"created_at"`
}
