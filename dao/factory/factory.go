package factory

import (
	"awesomeProject/dao/imports"
	"awesomeProject/dao/interfaces"
	"awesomeProject/dao/postgrsql"
)

func FactoryPersona() interfaces.PersonaDAO {
	var i interfaces.PersonaDAO
	i = postgrsql.PersonaImpl{}
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

func FactoryLog() interfaces.LogDAO {
	var i interfaces.LogDAO
	i = postgrsql.LogImpl{}
	return i
}

func FactoryLoginReg() interfaces.LoginRegDAO {
	var i interfaces.LoginRegDAO
	i = postgrsql.LoginReg{}
	return i
}

func FactoryRol() interfaces.RolDAO {
	var i interfaces.RolDAO
	i = postgrsql.RolImpl{}
	return i
}