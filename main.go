package main

import (
	"api-pdf/pdf"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// type Data struct {
// 	nombreSistema  string
// 	versionSistema string
// 	usuarioNombre  string
// 	celularAxeso   string
// 	sede           string
// 	cargo          string
// 	personaReporte string
// 	celularPersona string
// 	fecha          string
// 	descripcion    string
// 	descartes      string
// }

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/pdf", handlePDFRequestGin)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}

	// http.HandleFunc("/pdf", handlePDFRequest)
	// fmt.Println("Run server on http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePDFRequestGin(c *gin.Context) {

	var data pdf.Data

	if err := c.BindJSON(&data); err != nil {
		// c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pdfBytes, err := pdf.CrearPdf(&data)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pdfMemory, err := ioutil.ReadAll(pdfBytes)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Envía el flujo de bytes del PDF como respuesta
	c.Header("Content-Type", "application/pdf")
	// c.Header("Content-Disposition", "attachment; filename=pdf.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfMemory)
}

/*

func handlePDFRequest(w http.ResponseWriter, r *http.Request) {

	pdfBytes, err := pdf.CrearPdf()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Envía el flujo de bytes del PDF al cliente
	w.Header().Set("Content-Type", "application/pdf")
	// w.Header().Set("Content-Disposition", "attachment; filename=pdf.pdf")
	_, err = w.Write([]byte(pdfBytes.String()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
*/
