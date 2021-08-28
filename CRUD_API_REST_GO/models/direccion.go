package models

type Direccion struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_incremet"`
	Calle string `json:"calle"`
	Pais  string `json:"pais"`
}
