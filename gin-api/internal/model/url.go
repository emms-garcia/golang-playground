package model

type Url struct {
	ID        int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Original  string `json:"original" gorm:"not null"`
	ShortCode string `json:"short_code" gorm:"not null;unique"`
	Usages    int    `json:"usages" gorm:"default:0"`
}

func (u *Url) GetShortUrl() string {
	return "http://localhost/u/" + u.ShortCode
}
