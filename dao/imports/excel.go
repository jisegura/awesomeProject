package imports

import (
	"awesomeProject/dao/postgrsql"
	"awesomeProject/models"
	"github.com/tealeg/xlsx"
	"strconv"
)

type ExcelImpl struct{}

func newExcel() models.Excel {

	excel := models.Excel{}
	excel.Caja1 = []string{"Fecha inicio", "Fecha fin", "Inicio caja", "Cierre caja", "Cierre real", "Diferencia"}
	excel.Caja2 = []string{"Ingreso efectivo", "Ingreso crédito", "Ingreso débito", "Retiros", "Gastos", "Ingreso total"}
	excel.Facturas = []string{"Tipo", "Empleado", "Hora", "Precio", "Descuento", "Forma de pago", "Comentario", "Comentario baja"}
	excel.Renglones = []string{"Producto", "Cantidad", "Precio", "Descuento"}

	return excel
}

func (dao ExcelImpl) Export(movimientos []models.Movimientos) error {

	file := xlsx.NewFile()
	excel := newExcel()

	for i := range movimientos {
		sheet, err := file.AddSheet("Caja " + strconv.Itoa(i))
		if err != nil {
			return err
		}

		row := sheet.AddRow()
		cell := row.AddCell()
		cell.SetStyle(titleStyle())
		cell.Value = "Caja"
		row = sheet.AddRow()
		for j := range excel.Caja1 {
			cell = row.AddCell()
			cell.SetStyle(titleFactStyle())
			cell.Value = excel.Caja1[j]
		}

		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		cell.Value = movimientos[i].Caja.HoraInicio.Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		cell.Value = (movimientos[i].Caja.HoraFin).Format("2006-01-02 15:04:05")
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Caja inicio
		cell.Value = strconv.FormatFloat(movimientos[i].Caja.Inicio, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Cierre caja
		cell.Value = strconv.FormatFloat(movimientos[i].Caja.Fin, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Fin real
		cell.Value = strconv.FormatFloat(movimientos[i].Caja.CierreReal, 'f', 2, 64)
		cell = row.AddCell()
		//Diferencia
		diferencia := movimientos[i].Caja.CierreReal - movimientos[i].Caja.Fin
		if diferencia > 0 {
			cell.SetStyle(dataStyle())
		} else {
			cell.SetStyle(perdidaStyle())
		}

		cell.Value = strconv.FormatFloat(diferencia, 'f', 2, 64)
		row = sheet.AddRow()

		for j := range excel.Caja2 {
			cell = row.AddCell()
			cell.SetStyle(titleFactStyle())
			cell.Value = excel.Caja2[j]
		}
		row = sheet.AddRow()

		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Ingreso efectivo
		efectivo, err := postgrsql.Get_ingresos(movimientos[i].Caja.Id_caja, 1)
		if err != nil {
			return err
		}
		cell.Value = strconv.FormatFloat(efectivo, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Ingreso credito
		credito, err := postgrsql.Get_ingresos(movimientos[i].Caja.Id_caja, 3)
		if err != nil {
			return err
		}
		cell.Value = strconv.FormatFloat(credito, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Ingreso debito
		debito, err := postgrsql.Get_ingresos(movimientos[i].Caja.Id_caja, 2)
		if err != nil {
			return err
		}
		cell.Value = strconv.FormatFloat(debito, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		retiros, err := postgrsql.GetTotalRetiros(movimientos[i].Caja.Id_caja)
		if err != nil {
			return err
		}
		//Retiros
		cell.Value = strconv.FormatFloat(retiros, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		gastos, err := postgrsql.GetTotalGastos(movimientos[i].Caja.Id_caja)
		if err != nil {
			return err
		}
		//Gastos
		cell.Value = strconv.FormatFloat(gastos, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Ingreso total
		ingresoTotal := movimientos[i].Caja.Inicio + efectivo + credito + debito - retiros - gastos
		cell.Value = strconv.FormatFloat(ingresoTotal, 'f', 2, 64)
		row = sheet.AddRow()
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(titleStyle())
		cell.Value = "Facturas"

		for facturas := range movimientos[i].Facturas {

			row = sheet.AddRow()
			for j := range excel.Facturas {
				cell = row.AddCell()
				cell.SetStyle(titleFactStyle())
				cell.Value = excel.Facturas[j]
			}
			row = sheet.AddRow()

			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = postgrsql.GetTipo(movimientos[i].Facturas[facturas])
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			nombre, err := postgrsql.GetNombre(movimientos[i].Facturas[facturas].Id_empleado.Int64)
			if err != nil {
				return err
			}
			cell.Value = nombre
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = movimientos[i].Facturas[facturas].Fecha.Format("15:04:05")
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Precio, 'f', 2, 64)
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = strconv.FormatInt(movimientos[i].Facturas[facturas].Descuento.Int64, 10)
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = strconv.FormatInt(movimientos[i].Facturas[facturas].FormaDePago.Int64, 10)
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = (movimientos[i].Facturas[facturas].Comentario).String
			cell = row.AddCell()
			cell.SetStyle(dataStyle())
			cell.Value = movimientos[i].Facturas[facturas].ComentarioBaja
			cell = row.AddCell()
			row = sheet.AddRow()

			if len(movimientos[i].Facturas[facturas].Renglones) != 0 {
				row = sheet.AddRow()
				for j := range excel.Renglones {
					cell = row.AddCell()
					cell.SetStyle(titleRenglonStyle())
					cell.Value = excel.Renglones[j]

				}
				row = sheet.AddRow()
				for renglones := range movimientos[i].Facturas[facturas].Renglones {
					cell = row.AddCell()
					producto, err := postgrsql.GetNombreById(movimientos[i].Facturas[facturas].Renglones[renglones].Id_producto.Int64)
					if err != nil {
						return err
					}
					cell.SetStyle(dataStyle())
					cell.Value = producto
					cell = row.AddCell()
					cell.SetStyle(dataStyle())
					cell.Value = strconv.Itoa(movimientos[i].Facturas[facturas].Renglones[renglones].Cantidad)
					cell = row.AddCell()
					cell.SetStyle(dataStyle())
					cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Renglones[renglones].Precio, 'f', 2, 64)
					cell = row.AddCell()
					cell.SetStyle(dataStyle())
					cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Renglones[renglones].Descuento, 'f', 2, 64)
					row = sheet.AddRow()
				}
			}

		}
		row = sheet.AddRow()

	}

	err := file.Save("./Movimientos.xlsx")
	if err != nil {
		return err
	}
	return nil

}

func titleStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	font := *xlsx.NewFont(12, "Verdana")
	font.Bold = true
	style.Font = font
	fill := *xlsx.NewFill("solid", "76933c", "76933c")
	style.Fill = fill
	border := *xlsx.NewBorder("medium", "medium", "medium", "medium")
	border.BottomColor = "4f6228"
	border.TopColor = "4f6228"
	border.LeftColor = "4f6228"
	border.RightColor = "4f6228"
	style.Border = border
	style.ApplyFont = true
	style.ApplyFill = true
	style.ApplyBorder = true

	return style
}

func titleFactStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "c4d79b", "c4d79b")
	style.Fill = fill

	style.ApplyFill = true

	return style
}

func titleRenglonStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "d8e4bc", "d8e4bc")
	style.Fill = fill
	style.ApplyFill = true

	return style
}

func dataStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "ebf1de", "ebf1de")
	style.Fill = fill
	style.ApplyFill = true

	return style
}

func perdidaStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "ff0000", "ff0000")
	style.Fill = fill
	style.ApplyFill = true

	return style
}
