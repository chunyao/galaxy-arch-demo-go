package model

type ReportSubscribe struct {
	Id             int    `json:"id" gorm:"primary_key"`
	UnionId        string `json:"union_id" gorm:"size:255"`
	MabangAccount  string `json:"mabang_account" gorm:"size:255"`
	MabangPassword string `json:"mabang_password" gorm:"size:255"`
	PageType       string `json:"page_type" gorm:"size:50"`
	PageUrl        string `json:"page_url"`
	CosUrl         string `json:"cos_url"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
	JobDate        string `json:"cos_url"`
}
