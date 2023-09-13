package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PersonResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (p Person) ToResponse() *PersonResponse {
	return &PersonResponse{
		ID:   p.ID,
		Name: p.Name,
	}
}
