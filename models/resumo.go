package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Resumo struct {
	gorm.Model
	TotalReceitas    int `json:"totalreceitas" validate:"regexp=^[0-9]$"`
	TotalDespesas    int `json:"totaldespesas" validate:"regexp=^[0-9]$"`
	Saldo            int `json:"saldo" validate:"regexp=^[0-9]$"`
	GastoAlimentacao int `json:"gasto_alimentacao" validate:"regexp=^[0-9]$"`
	GastoSaude       int `json:"gasto_saude" validate:"regexp=^[0-9]$"`
	GastoMoradia     int `json:"gasto_moradia" validate:"regexp=^[0-9]$"`
	GastoTransporte  int `json:"gasto_transporte" validate:"regexp=^[0-9]$"`
	GastoEducacao    int `json:"gasto_educacao" validate:"regexp=^[0-9]$"`
	GastoLazer       int `json:"gasto_lazer" validate:"regexp=^[0-9]$"`
	GastoImprevistos int `json:"gasto_imprevistos" validate:"regexp=^[0-9]$"`
	GastoOutras      int `json:"gasto_outras" validate:"regexp=^[0-9]$"`
}

func ValidaDadosResumo(resumo *Resumo) error {
	if err := validator.Validate(resumo); err != nil {
		return err
	}
	return nil

}
