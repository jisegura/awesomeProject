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
	excel.Caja2 = []string{"Ingreso efectivo", "Ingreso crédito", "Ingreso débito", "Retiros", "Gastos", "Ingreso total", "Cierre fiscal"}
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
			cell.SetStyle(titleCajaStyle())
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
			cell.SetStyle(titleCajaStyle())
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
		ingresoTotal := efectivo + credito + debito
		cell.Value = strconv.FormatFloat(ingresoTotal, 'f', 2, 64)
		cell = row.AddCell()
		cell.SetStyle(dataStyle())
		//Cierre fiscal
		cell.Value = strconv.FormatFloat(movimientos[i].Caja.CierreFiscal, 'f', 2, 64)

		row = sheet.AddRow()
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(titleStyle())
		cell.Value = "Facturas"

		//FACTURAS

		for facturas := range movimientos[i].Facturas {

			style := dataStyle()
			row = sheet.AddRow()
			for j := range excel.Facturas {
				cell = row.AddCell()
				if j == len(excel.Facturas)-1 {
					cell.SetStyle(borderTitleStyle())
				} else {
					cell.SetStyle(titleFacturaStyle())
				}
				cell.Value = excel.Facturas[j]
			}
			row = sheet.AddRow()

			cell = row.AddCell()
			//Tipo
			cell.SetStyle(style)
			cell.Value = postgrsql.GetTipo(movimientos[i].Facturas[facturas])
			cell = row.AddCell()
			//Nombre empleado
			cell.SetStyle(style)
			nombre, err := postgrsql.GetNombre(movimientos[i].Facturas[facturas].Id_empleado.Int64)
			if err != nil {
				return err
			}
			cell.Value = nombre
			cell = row.AddCell()
			//Fecha
			cell.SetStyle(style)
			cell.Value = movimientos[i].Facturas[facturas].Fecha.Format("15:04:05")
			cell = row.AddCell()
			cell.SetStyle(style)
			//Precio
			cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Precio, 'f', 2, 64)
			cell = row.AddCell()
			//Descuento
			if movimientos[i].Facturas[facturas].Descuento.Int64 != 0 {
				cell.SetStyle(descuentoStyle())
			} else {
				cell.SetStyle(style)
			}
			cell.Value = strconv.FormatInt(movimientos[i].Facturas[facturas].Descuento.Int64, 10)
			cell = row.AddCell()
			cell.SetStyle(style)
			//Forma de pago
			if movimientos[i].Facturas[facturas].FormaDePago.Valid {
				cell.Value = postgrsql.GetFormaDePago(movimientos[i].Facturas[facturas])
			}
			cell = row.AddCell()
			cell.SetStyle(style)
			//Comentario
			cell.Value = (movimientos[i].Facturas[facturas].Comentario).String
			cell = row.AddCell()
			//Comentario baja
			if len(movimientos[i].Facturas[facturas].ComentarioBaja) != 0 {
				cell.SetStyle(comentarioStyle())
			} else {
				cell.SetStyle(borderDataStyle())
			}
			cell.Value = movimientos[i].Facturas[facturas].ComentarioBaja
			row = sheet.AddRow()

			if len(movimientos[i].Facturas[facturas].Renglones) == 0 {
				for range excel.Facturas {
					cell = row.AddCell()
					cell.SetStyle(lastDataStyle())
				}
			} else {
				for j := range excel.Facturas {
					cell = row.AddCell()
					if j == len(excel.Facturas)-1 {
						cell.SetStyle(borderDataStyle())
					} else {
						cell.SetStyle(dataStyle())
					}
				}
				row = sheet.AddRow()
				for j := range excel.Facturas {
					cell = row.AddCell()
					if j == len(excel.Facturas)-1 {
						cell.SetStyle(borderRenglonStyle())
					} else {
						cell.SetStyle(titleRenglonStyle())
					}
					if j < len(excel.Renglones) {
						cell.Value = excel.Renglones[j]
					}
				}

				row = sheet.AddRow()
				for renglones := range movimientos[i].Facturas[facturas].Renglones {

					style = dataStyle()

					cell = row.AddCell()
					producto, err := postgrsql.GetNombreById(movimientos[i].Facturas[facturas].Renglones[renglones].Id_producto.Int64)
					if err != nil {
						return err
					}

					cell.SetStyle(style)
					cell.Value = producto
					cell = row.AddCell()
					cell.SetStyle(style)
					cell.Value = strconv.Itoa(movimientos[i].Facturas[facturas].Renglones[renglones].Cantidad)
					cell = row.AddCell()
					cell.SetStyle(style)
					cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Renglones[renglones].Precio, 'f', 2, 64)
					cell = row.AddCell()
					//Descuento
					if movimientos[i].Facturas[facturas].Renglones[renglones].Descuento != 0 {
						cell.SetStyle(descuentoStyle())
					} else {
						cell.SetStyle(style)
					}
					cell.Value = strconv.FormatFloat(movimientos[i].Facturas[facturas].Renglones[renglones].Descuento, 'f', 2, 64)
					for col := len(excel.Renglones); col < len(excel.Facturas); col++ {
						cell = row.AddCell()
						cell.SetStyle(style)
						if col == len(excel.Facturas)-1 {
							cell.SetStyle(borderDataStyle())
						}
					}

					row = sheet.AddRow()
				}
				for range excel.Facturas {
					cell = row.AddCell()
					cell.SetStyle(lastDataStyle())
				}
			}

		}
		row = sheet.AddRow()
	}

	err := file.Save("./" + movimientos[0].Caja.HoraInicio.Format("2006-01-02") + "to" + movimientos[len(movimientos)-1].Caja.HoraInicio.Format("2006-01-02") + ".xlsx")
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

func titleFacturaStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "solid", "medium", "solid")
	fill := *xlsx.NewFill("solid", "c4d79b", "c4d79b")
	style.Border = border
	style.Fill = fill

	style.ApplyBorder = true
	style.ApplyFill = true
	return style
}

func titleCajaStyle() *xlsx.Style {

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

func lastDataStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "solid", "medium", "solid")
	style.Border = border

	style.ApplyBorder = true
	return style
}

func borderTitleStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "medium", "medium", "solid")
	fill := *xlsx.NewFill("solid", "c4d79b", "c4d79b")
	style.Border = border
	style.Fill = fill

	style.ApplyBorder = true
	style.ApplyFill = true

	return style
}

func borderDataStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "medium", "solid", "solid")
	fill := *xlsx.NewFill("solid", "ebf1de", "ebf1de")

	style.Border = border
	style.Fill = fill
	style.ApplyBorder = true
	style.ApplyFill = true

	return style
}

func borderRenglonStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "medium", "solid", "solid")
	fill := *xlsx.NewFill("solid", "d8e4bc", "d8e4bc")

	style.Border = border
	style.Fill = fill
	style.ApplyBorder = true
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

func descuentoStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	fill := *xlsx.NewFill("solid", "ffc000", "ffc000")
	style.Fill = fill
	style.ApplyFill = true

	return style
}

func comentarioStyle() *xlsx.Style {

	style := xlsx.NewStyle()
	border := *xlsx.NewBorder("solid", "medium", "solid", "solid")
	fill := *xlsx.NewFill("solid", "ffff00", "ffff00")
	style.Border = border
	style.Fill = fill
	style.ApplyBorder = true
	style.ApplyFill = true

	return style
}
