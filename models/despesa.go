package models

import (
	"strings"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Despesa struct {
	gorm.Model
	Descricao string  `json:"descricao" validate:"nonzero"`
	Valor     float32 `json:"valor" validate:"min=0.1"`
	Data      string  `json:"data" validate:"len=7, regexp=^[0-9]{2}/[0-9]{4}$"`
	Categoria string  `json:"categoria"`
}

func ValidaDadosDespesa(despesa *Despesa) error {
	if err := validator.Validate(despesa); err != nil {
		return err
	}
	return nil
}

func ValidaCategoria(despesa *Despesa) string {
	var esperado = []string{"alimentação", "saúde", "moradia", "transporte", "educação", "lazer", "imprevistos", "outras"}
	minusculas := strings.ToLower(despesa.Categoria)
	retorno := ""
	for i := 0; i < len(esperado); i++ {
		if minusculas == esperado[i] {
			retorno = esperado[i]
			return retorno
		} else {
			retorno = "outras"
		}
	}
	return retorno
}
