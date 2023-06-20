package pdf

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"

	"github.com/johnfercher/maroto/pkg/props"
)

func CrearPdf(pdfData Data) (*bytes.Buffer, error) {

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)

	builHeading(m)
	builBody(m, pdfData)

	bufferPdf, err := m.Output()
	if err != nil {
		return nil, err
	}

	return &bufferPdf, nil
}

func builHeading(m pdf.Maroto) {

	m.RegisterHeader(func() {

		m.SetBorder(true)

		m.Row(25, func() {
			m.Col(3, func() {
				err := m.FileImage("pdf/logo.png", props.Rect{
					Center:  true,
					Percent: 75,
				})

				if err != nil {
					m.Text("Logo", props.Text{
						Top:  1,
						Size: 12,
						// Style: consts.Bold,
						Align: consts.Center,
					})
				}
			})

			m.Col(6, func() {
				m.Text("PODER JUDICIAL", props.Text{
					Top:   5,
					Style: consts.Bold,
					Align: consts.Center,
					Size:  16,
				})
				m.Text("FORMULARIO DE INCIDENCIA", props.Text{
					Top:   15,
					Align: consts.Center,
					Size:  13,
				})
			})

			m.Col(3, func() {
				m.Text("Formulario N°:", props.Text{
					Top:   1,
					Style: consts.Bold,
					Align: consts.Center,
					// Color: darkPurpleColor(),
				})
				m.Text("P-01-2023-GI-GG-PJ-F-01", props.Text{
					Top:   7,
					Align: consts.Center,
				})

				m.Text("Versión: ", props.Text{
					Top:   13,
					Style: consts.Bold,
					Right: 10,
					Align: consts.Middle,
				})
				m.Text("05", props.Text{
					Top:   13,
					Left:  10,
					Align: consts.Middle,
				})

				m.Text("Fecha: ", props.Text{
					Top:   19,
					Style: consts.Bold,
					Right: 15,
					Align: consts.Middle,
				})
				m.Text("08/06/2023", props.Text{
					Top:   19,
					Left:  15,
					Align: consts.Middle,
				})

			})

		})

		m.SetBorder(false)

		m.Line(8, props.Line{
			Color: color.NewWhite(),
		})

		m.SetBorder(true)

	})

}

func builBody(m pdf.Maroto, info Data) {

	// INFORMACIÓN DEL SISTEMA
	m.SetBorder(true)

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INFORMACIÓN DEL SISTEMA", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombre", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.NombreSistema, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Versión", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.VersionSistema, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// INFORMACIÓN DEL USUARIO

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INFORMACIÓN DEL USUARIO", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombres y Apellidos", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.UsuarioNombre, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Celular / Anexo", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.CelularAxeso, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})
	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Sede", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.Sede, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Cargo", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.Cargo, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// REPORTADO POR

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("REPORTADO POR", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(3, func() {
			m.Text("Nombres y Apellidos", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.PersonaReporte, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text("Celular", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
		m.Col(3, func() {
			m.Text(info.CelularPersona, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
	})

	// INCIDENCIA

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("INCIDENCIA", props.Text{
				Top:   1,
				Size:  12,
				Right: 25,
				Style: consts.Bold,
				Align: consts.Middle,
			})
			m.Text(info.Fecha, props.Text{
				Top:   1,
				Size:  12,
				Left:  25,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(info.Descripcion, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})

	})

	// DESCARTES

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("DESCARTES", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())
	fmt.Println("Si o no: " + info.DescarteAcepta)
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text(info.Descartes+" "+info.DescarteAcepta, props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})

	})

	// FLUJO

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("FLUJO", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Middle,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Pantalla_01 (Obligatorio SIJ): Versión del sistema | Usuario | Cargo ", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})

	})

	var count int
	for _, item := range info.Imagenes {
		count++
		m.Row(75, func() {
			m.Col(12, func() {
				err := m.FileImage("pdf/"+strconv.Itoa(count)+"output."+item.Extension, props.Rect{
					Center:  true,
					Percent: 95,
				})

				if err != nil {
					m.Text("Captura", props.Text{
						Top:  1,
						Size: 12,
						// Style: consts.Bold,
						Align: consts.Center,
					})
				}
			})
		})

		m.SetBackgroundColor(color.NewWhite())

		m.Row(20, func() {
			m.Col(12, func() {
				m.Text(item.Descripcion, props.Text{
					Extrapolate: false,
					Top:         1,
					Left:        1,
					Size:        10,
					Style:       consts.Bold,
					Align:       consts.Left,
				})

			})

		})
	}

	// Parámetros de prueba de Desarrollo*

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Parámetros de prueba de Desarrollo*", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(20, func() {
		m.Col(6, func() {
			m.Text("[Nombre del Servicio**]:", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Nombre BD]:", props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[User Id]:", props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(6, func() {
			m.Text("[Passwd]:", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[IP]:", props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Puerto]:", props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})

	})

	// Parámetros de prueba de Calidad (Testing)*

	m.SetBackgroundColor(grayColor())
	m.Row(8, func() {
		m.Col(12, func() {
			m.Text("Parámetros de prueba de Calidad (Testing)*", props.Text{
				Top:   1,
				Size:  12,
				Style: consts.Bold,
				Align: consts.Left,
			})
		})
	})
	m.SetBackgroundColor(color.NewWhite())

	m.Row(20, func() {
		m.Col(6, func() {
			m.Text("[Nombre del Servicio**]:", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Nombre BD]:", props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[User Id]:", props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})
		m.Col(6, func() {
			m.Text("[Passwd]:", props.Text{
				Top:   1,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[IP]:", props.Text{
				Top:   7,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
			m.Text("[Puerto]:", props.Text{
				Top:   13,
				Left:  1,
				Size:  10,
				Style: consts.Normal,
				Align: consts.Left,
			})
		})

	})

	m.SetBorder(false)

}

func grayColor() color.Color {
	return color.Color{
		Red:   220,
		Green: 220,
		Blue:  220,
	}
}
