package model

// Url struct represents a URL entity in the database.
type Url struct {
	ID        int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Original  string `json:"original" gorm:"not null"`
	ShortCode string `json:"short_code" gorm:"not null;unique"`
	Usages    int    `json:"usages" gorm:"default:0"`
}
