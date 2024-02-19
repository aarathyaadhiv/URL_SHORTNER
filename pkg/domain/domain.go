package domain

type Url struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Original string `json:"original"`
	Shorten  string `json:"shorten"`
	Counts    int    `json:"count" gorm:"default:0"`
}
