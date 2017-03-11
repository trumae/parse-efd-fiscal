package NotaFiscal

import "github.com/jinzhu/gorm"

type Item struct {
	gorm.Model
	Codigo string		`gorm:"type:varchar(60)"`
	Ean string		`gorm:"type:varchar(60)"`
	Descricao string	`gorm:"type:varchar(100)"`
	Ncm string		`gorm:"type:varchar(60)"`
	Cfop string		`gorm:"type:varchar(60)"`
	Unid string		`gorm:"type:varchar(60)"`
	Qtd float64		`gorm:"type:decimal(19,3)"`
	VUnit float64		`gorm:"type:decimal(19,3)"`
	VTotal float64		`gorm:"type:decimal(19,3)"`
	NotaFiscalID uint
}