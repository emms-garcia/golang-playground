package model

type Url struct {
	ID        int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Original  string `json:"url" gorm:"not null"`
	ShortCode string `json:"short_code" gorm:"not null;unique"`
}

func (u *Url) GetShortUrl() string {
	return "http://localhost/u/" + u.ShortCode
}
