package models

type Persona struct {
	// gorm.Model
	ID int64 `json:"id" gorm:"primary_key;auto_incremet"`
	//ID        int64  `json:"id"`
	Nombre      string    `json:"nombre"`
	Apellido    string    `json:"apellido"`
	Lugar       string    `json:"lugar"`
	Telefono    string    `json:"telefono"`
	DireccionID int64     `json:"direccionId"`
	Direcciones Direccion `gorm:"foreignKey:DireccionID"`
	// Direcciones Direccion `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
