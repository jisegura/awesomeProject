package interfaces

import "awesomeProject/models"

type ExcelDAO interface {
	Export(movimientos []models.Movimientos) error
}
