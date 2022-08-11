package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Receita struct {
	gorm.Model
	Descricao string  `json:"descricao" validate:"nonzero"`
	Valor     float32 `json:"valor" validate:"min=0.1"`
	Data      string  `json:"data" validate:"nonzero"`
}

func ValidaDadosReceita(receita *Receita) error {
	if err := validator.Validate(receita); err != nil {
		return err
	}
	return nil

}
