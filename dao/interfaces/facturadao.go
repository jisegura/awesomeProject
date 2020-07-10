package interfaces

import (
	"awesomeProject/models"
)

type FacturaDAO interface {

	Create(factura *models.Factura) error
	Create_Customer(factura *models.Factura) (models.Factura, error)
	Create_Other(factura *models.Factura) (models.Factura, error)
	Get_All(id int) ([]models.Factura, error)
	Get_Customer_Invoices(id int) ([]models.Factura, error)
	Get_Other_Invoices(id int) ([]models.Factura, error)
	Get_By_Id(id int) (models.Factura, error)
	Get_By_Customer_Id(id int) (models.Factura, error)
	Get_By_Other_Id(id int) (models.Factura, error)
	Update_Comment(factura *models.Factura) error
	Get_Deleted_Invoices() ([]models.Factura, error)
	Get_Invoice_Id_By_Cash_Box(id int) ([]int, error)
	Get_All_Ivloices_By_Cash_Box(id int) ([]models.Factura, error)
	Get_Last_Invoices(id int) ([]models.Factura, error)
	Get_Incomes(id int, formaDePago int) (float64, error)
}
