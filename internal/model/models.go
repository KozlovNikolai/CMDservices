package model

type Service struct {
	ID    uint64 `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int64  `json:"price" db:"price"`
}
