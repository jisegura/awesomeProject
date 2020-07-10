package interfaces

import (
	"awesomeProject/models"
	"time"
)

type CajaDAO interface {
	Create(caja *models.Caja) (models.Caja, error)
	Get_Cash_Box() (models.Caja, error)
	Get_By_Id(id int) (models.Caja, error)
	Get_All() ([]models.Caja, error)
	Close_Cash_Box (cashBox *models.Caja) (models.Caja, error)
	Get_By_Date(startDate time.Time, finalDate time.Time) ([]models.Caja, error)
}
