package factory

import (
	"awesomeProject/dao/imports"
	"awesomeProject/dao/interfaces"
	"awesomeProject/dao/postgrsql"
)

func FactoryEmpleado() interfaces.EmpleadoDAO {
	var i interfaces.EmpleadoDAO
	i = postgrsql.EmpleadoImpl{}
	return i
}

func FactoryProducto() interfaces.ProductoDAO {
	var i interfaces.ProductoDAO
	i = postgrsql.ProductoImpl{}
	return i
}

func FactoryCaja() interfaces.CajaDAO {
	var i interfaces.CajaDAO
	i = postgrsql.CajaImpl{}
	return i
}

func FactoryFactura() interfaces.FacturaDAO {
	var i interfaces.FacturaDAO
	i = postgrsql.FacturaImpl{}
	return i
}

func FactoryRenglon() interfaces.RenglonDAO {
	var i interfaces.RenglonDAO
	i = postgrsql.RenglonImpl{}
	return i
}

func FactoryCategoria() interfaces.CategoriaDAO {
	var i interfaces.CategoriaDAO
	i = postgrsql.CategoriaImpl{}
	return i
}

func FactoryExcel() interfaces.ExcelDAO {
	var i interfaces.ExcelDAO
	i = imports.ExcelImpl{}
	return i
}

func FactoryUser() interfaces.UserDAO {
	var i interfaces.UserDAO
	i = postgrsql.UserImpl{}
	return i
}
