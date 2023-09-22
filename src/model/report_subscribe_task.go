package model

type ReportSubscribeTask struct {
	Id             int    `json:"id" gorm:"primary_key"`
	UnionId        string `json:"union_id" gorm:"size:255"`
	MabangAccount  string `json:"mabang_account" gorm:"size:255"`
	MabangPassword string `json:"mabang_password" gorm:"size:255"`
	CosUrl         string `json:"cos_url"`
	JobDate        string `json:"job_date"`
	Status         int    `json:"status" gorm:"size:1"`
}
