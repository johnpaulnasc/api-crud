package domain

type Item struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}