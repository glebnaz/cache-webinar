package model

type Post struct {
	Body     string `json:"body"`
	CreateBy string `json:"create_by"`
	CreateAt int64  `json:"create_at"`
}
