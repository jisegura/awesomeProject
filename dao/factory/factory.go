package factory

import (
"awesomeProject/dao/interfaces"
"awesomeProject/dao/postgrsql"
)

func FactoryUser ()interfaces.UserDAO {
	var i interfaces.UserDAO
	i = postgrsql.UserImpl{}
	return i
}

func FactoryProducto ()interfaces.ProductoDAO {
	var i interfaces.ProductoDAO
	i = postgrsql.ProductoImpl{}
	return i
}

func FactoryCaja ()interfaces.CajaDAO {
	var i interfaces.CajaDAO
	i = postgrsql.CajaImpl{}
	return i
}

func FactoryRenglon ()interfaces.RenglonDAO {
	var i interfaces.RenglonDAO
	i = postgrsql.RenglonImpl{}
	return i
}

func FactoryFactura ()interfaces.FacturaDAO {
	var i interfaces.FacturaDAO
	i = postgrsql.FacturaImpl{}
	return i
}